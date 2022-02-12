package internal

import (
	"errors"
	"strings"
)

func NewGame(secret string) (*Game, error) {
	if secret == "" {
		return nil, errors.New("secret must be a 5-letter word")
	}
	return &Game{secret: strings.ToUpper(secret)}, nil
}

type Game struct {
	secret string
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

func (g *Game) Guess(guess string) GuessResult {
	guess = strings.ToUpper(guess)

	res := GuessResult{}
	for i := 0; i < 5; i++ {
		clue := Clue{
			Letter: rune(guess[i]),
			Result: Wrong,
		}
		if guess[i] == g.secret[i] {
			clue.Result = Correct
		} else if strings.ContainsRune(g.secret, rune(guess[i])) {
			clue.Result = Misplaced
		} else {
			clue.Result = Wrong
		}
		res.Clues[i] = clue
	}
	return res
}
