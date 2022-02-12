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
		for j, c := range res.Clues {
			if c.Result != Wrong {
				continue
			}
			if c.Letter == rune(g.secret[i]) {
				res.Clues[j].Result = Misplaced
				continue
			}
		}
	}

	//for i := 0; i < 5; i++ {
	//	clue := Clue{
	//		Letter: rune(guess[i]),
	//		Result: Wrong,
	//	}
	//	if guess[i] == g.secret[i] {
	//		clue.Result = Correct
	//	} else if strings.ContainsRune(g.secret, rune(guess[i])) {
	//		clue.Result = Misplaced
	//	} else {
	//		clue.Result = Wrong
	//	}
	//	res.Clues[i] = clue
	//}
	return res, nil
}
