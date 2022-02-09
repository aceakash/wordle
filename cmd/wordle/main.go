package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	game, err := NewGame()
	if err != nil {
		panic(err)
	}
	ui := NewUI(os.Stdin, os.Stdout, os.Stderr)
	ui.MessageUser("You're guessing a 5-letter word.")

	solved := false
	guessCount := 0
	for guessCount < 6 {
		err = ui.RenderGameState(game)
		if err != nil {
			panic(err)
		}
		guess, err := ui.GetGuessFromUser()
		if !game.isGuessValid(guess) {
			ui.InvalidGuess()
			continue
		}
		//fmt.Println("guess", guess)
		if err != nil {
			panic(err)
		}
		//fmt.Println("===", guess, game.Secret(), guess == game.Secret(), "===")
		game.RegisterGuess(guess)
		if guess == game.Secret() {
			solved = true
			break
		}
		guessCount++
	}
	ui.RenderGameState(game)
	if !solved {
		ui.NotSolved(game.Secret())
	} else {
		ui.GameSolved()
	}
}

type UI struct {
	out io.Writer
	in  io.Reader
	err io.Writer
}

func (ui UI) RenderGameState(game *Game) error {
	_, err := fmt.Fprintln(ui.out, "==========")
	if err != nil {
		return err
	}
	//_, err = fmt.Fprintf(ui.out, "The secret is %s\n", game.Secret())
	//if err != nil {
	//	return err
	//}
	green := color.New(color.BgGreen)
	yellow := color.New(color.BgYellow)
	black := color.New(color.BgHiBlack)
	for _, c := range game.Clues() {
		for _, r := range c {
			if r.Type == SpotOn {
				green.Fprintf(ui.out, "%c ", r.Letter)
				continue
			}
			if r.Type == WrongPlace {
				yellow.Fprintf(ui.out, "%c ", r.Letter)
				continue
			}
			black.Fprintf(ui.out, "%c ", r.Letter)
		}
		fmt.Fprintln(ui.out, "")
	}
	return nil
}

func (ui UI) GetGuessFromUser() (string, error) {
	reader := bufio.NewReader(ui.in)
	_, err := fmt.Fprintf(ui.out, "Enter guess: ")
	if err != nil {
		return "", err
	}
	guess, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(guess), nil
}

func (ui UI) GameSolved() error {
	_, err := fmt.Fprintln(ui.out, "Well done, that's correct!")
	if err != nil {
		return err
	}
	return nil
}

func (ui UI) NotSolved(secret string) {
	fmt.Fprintf(ui.out, "Unlucky, the word was %s\n", secret)
}

func (ui UI) InvalidGuess() {
	fmt.Fprintln(ui.out, "Uh oh, not one of the allowed 5-letter words!")
}

func (ui *UI) MessageUser(msg string) {
	fmt.Fprintln(ui.out, msg)
}

func NewUI(in io.Reader, out, err io.Writer) UI {
	return UI{out, in, err}
}

type Game struct {
	secret         string
	guesses        []string
	clues          []Clue
	allowedGuesses []string
}

func (g *Game) Secret() string {
	return g.secret
}

func (g *Game) RegisterGuess(guess string) {
	g.guesses = append(g.guesses, guess)
	clue := g.generateClue(guess)
	g.clues = append(g.clues, clue)
}

type ClueLetterType string

const (
	SpotOn     ClueLetterType = "SpotOn"
	WrongPlace ClueLetterType = "WrongPlace"
	Wrong      ClueLetterType = "Wrong"
)

type ClueLetter struct {
	Type   ClueLetterType
	Letter rune
}

type Clue []ClueLetter

func (g *Game) Clues() []Clue {
	return g.clues
}

func (g *Game) generateClue(guess string) Clue {
	clue := []ClueLetter{}
	for i, _ := range guess {
		if g.secret[i] == guess[i] {
			clue = append(clue, ClueLetter{
				Type:   SpotOn,
				Letter: rune(guess[i]),
			})
			continue
		}
		if strings.ContainsRune(g.secret, rune(guess[i])) {
			clue = append(clue, ClueLetter{
				Type:   WrongPlace,
				Letter: rune(guess[i]),
			})
			continue
		}
		clue = append(clue, ClueLetter{
			Type:   Wrong,
			Letter: rune(guess[i]),
		})
	}
	return clue
}

func (g *Game) isGuessValid(guess string) bool {
	return len(guess) == 5 && g.isAllowedGuess(guess)
}

func (g *Game) isAllowedGuess(guess string) bool {
	for _, allowedGuess := range g.allowedGuesses {
		//fmt.Printf(`"%s", "%s"\n`, guess, allowedGuess)
		if guess == allowedGuess {
			return true
		}
	}
	return false
}

func NewGame() (*Game, error) {
	secret, err := pickSecret()
	if err != nil {
		return &Game{}, err
	}
	allowedGuesses, err := getAllowedGuesses()
	if err != nil {
		return &Game{}, err
	}
	game := &Game{
		secret,
		[]string{},
		[]Clue{},
		allowedGuesses,
	}
	return game, nil
}

func getAllowedGuesses() ([]string, error) {
	bytes, err := os.ReadFile("wordle-allowed-guesses.txt")
	if err != nil {
		return nil, err
	}
	allowedRaw := strings.Split(string(bytes), "\n")

	bytes2, err := os.ReadFile("wordle-answers-alphabetical.txt")
	if err != nil {
		return nil, err
	}
	answersRaw := strings.Split(string(bytes2), "\n")

	words3 := []string{}
	words3 = append(words3, allowedRaw...)
	words3 = append(words3, answersRaw...)

	fmt.Println(contains(words3, "cater"))
	return words3, nil
}

func pickSecret() (string, error) {
	rand.Seed(time.Now().UnixNano())
	bytes, err := os.ReadFile("wordle-answers-alphabetical.txt")
	if err != nil {
		return "", err
	}
	words := strings.Split(string(bytes), "\n")
	return words[rand.Intn(len(words))], nil
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
