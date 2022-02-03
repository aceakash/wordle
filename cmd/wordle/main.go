package main

import (
	"bufio"
	"fmt"
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
	err = ui.RenderGameState(game)
	if err != nil {
		panic(err)
	}
	guess, err := ui.GetGuessFromUser()
	//fmt.Println("guess", guess)
	if err != nil {
		panic(err)
	}
	//fmt.Println("===", guess, game.Secret(), guess == game.Secret(), "===")
	if guess == game.Secret() {
		//fmt.Println("in")
		err := ui.GameSolved()
		if err != nil {
			panic(err)
		}
		return
	}
}

type UI struct {
	out io.Writer
	in  io.Reader
	err io.Writer
}

func (ui UI) RenderGameState(game Game) error {
	_, err := fmt.Fprintln(ui.out, "==========")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(ui.out, "The secret is %s\n", game.Secret())
	if err != nil {
		return err
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

func NewUI(in io.Reader, out, err io.Writer) UI {
	return UI{out, in, err}
}

type Game struct {
	secret string
}

func (g Game) Secret() string {
	return g.secret
}

func NewGame() (Game, error) {
	secret, err := pickSecret()
	if err != nil {
		return Game{}, err
	}
	game := Game{
		secret,
	}
	return game, nil
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
