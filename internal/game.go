package internal

import (
	"errors"
	"fmt"
	"strings"
)

type AllowedGuessChecker interface {
	IsGuessAllowed(guess string) bool
}

func NewGame(secret string, allowedGuessChecker AllowedGuessChecker) (*Game, error) {
	if secret == "" {
		return nil, errors.New("secret must be a 5-letter word")
	}
	return &Game{secret: strings.ToUpper(secret), guessChecker: allowedGuessChecker}, nil
}

type Game struct {
	secret       string
	guessChecker AllowedGuessChecker
}

type ClueResult string

const (
	Correct   ClueResult = "Correct"
	Wrong     ClueResult = "Wrong"
	Misplaced ClueResult = "Misplaced"
)

type Clue struct {
	Letter rune
	Result ClueResult
}

func (c Clue) LetterStr() string {
	return string(rune(c.Letter))
}

type GuessResult struct {
	Clues [5]Clue
}

func (gr GuessResult) IsCorrectAnswer() bool {
	for _, clue := range gr.Clues {
		if !(clue.Result == Correct) {
			return false
		}
	}
	return true
}

func (g *Game) Guess(guess string) (GuessResult, error) {
	guess = strings.ToUpper(guess)
	if !g.guessChecker.IsGuessAllowed(guess) {
		return GuessResult{}, fmt.Errorf("%s is not a 5-letter word known to me", guess)
	}

	res := GuessResult{}

	for i := 0; i < 5; i++ {
		res.Clues[i] = Clue{
			Letter: rune(guess[i]),
			Result: Wrong,
		}
	}

	for i := 0; i < 5; i++ {
		res.Clues[i].Letter = rune(guess[i])
		if guess[i] == g.secret[i] {
			res.Clues[i].Result = Correct
			continue
		}
		for j := 0; j < 5; j++ {
			if res.Clues[j].Result != Wrong {
				continue
			}
			if res.Clues[j].Letter == rune(g.secret[i]) {
				res.Clues[j].Result = Misplaced
				break
			}
		}
	}

	return res, nil
}
