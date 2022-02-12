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
	Correct ClueResult = "Correct"
)

type Clue struct {
	Letter string
	Result ClueResult
}

type GuessResult struct {
	Solved bool
	Clues  [5]Clue
}

func (g *Game) Guess(guess string) GuessResult {
	guess = strings.ToUpper(guess)
	if guess == g.secret {
		return GuessResult{
			Solved: true,
		}
	}
	if guess[0] == g.secret[0] {
		return GuessResult{
			Solved: false,
			Clues: [5]Clue{
				{
					Letter: string(g.secret[0]),
					Result: Correct,
				},
			},
		}
	}
	return GuessResult{}
}
