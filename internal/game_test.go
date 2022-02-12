package internal_test

import (
	i "github.com/aceakash/wordle/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	t.Run("returns error when no secret provided", func(t *testing.T) {
		_, err := i.NewGame("")
		assert.Error(t, err)
	})
}

func TestGame_WhenOnlyFirstLetterIsCorrect(t *testing.T) {
	g, _ := i.NewGame("folds")
	res := g.Guess("freak")

	expectedClues := [5]i.Clue{i.Clue{
		Letter: 'F',
		Result: i.Correct,
	},
		i.Clue{
			Letter: 'R',
			Result: i.Wrong,
		},
		i.Clue{
			Letter: 'E',
			Result: i.Wrong,
		},
		i.Clue{
			Letter: 'A',
			Result: i.Wrong,
		},
		i.Clue{
			Letter: 'K',
			Result: i.Wrong,
		},
	}
	assert.Equal(t, i.GuessResult{Solved: false, Clues: expectedClues}, res)
}

func Test_WhenGuessIsCorrectOnFirstTry(t *testing.T) {
	g, _ := i.NewGame("folds")
	res := g.Guess("folds")

	expectedClues := [5]i.Clue{i.Clue{
		Letter: 'F',
		Result: i.Correct,
	},
		i.Clue{
			Letter: 'O',
			Result: i.Correct,
		},
		i.Clue{
			Letter: 'L',
			Result: i.Correct,
		},
		i.Clue{
			Letter: 'D',
			Result: i.Correct,
		},
		i.Clue{
			Letter: 'S',
			Result: i.Correct,
		},
	}
	assert.Equal(t, i.GuessResult{Solved: false, Clues: expectedClues}, res)

}
