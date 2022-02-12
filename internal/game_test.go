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

func TestGame_CorrectLettersAreMarked(t *testing.T) {
	t.Run("first letter only", func(t *testing.T) {
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
	})

	t.Run("third letter only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("ruler")

		expectedClues := [5]i.Clue{i.Clue{
			Letter: 'R',
			Result: i.Wrong,
		},
			i.Clue{
				Letter: 'U',
				Result: i.Wrong,
			},
			i.Clue{
				Letter: 'L',
				Result: i.Correct,
			},
			i.Clue{
				Letter: 'E',
				Result: i.Wrong,
			},
			i.Clue{
				Letter: 'R',
				Result: i.Wrong,
			},
		}
		assert.Equal(t, i.GuessResult{Solved: false, Clues: expectedClues}, res)
	})

	t.Run("fifth letter only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("runes")

		expectedClues := [5]i.Clue{i.Clue{
			Letter: 'R',
			Result: i.Wrong,
		},
			i.Clue{
				Letter: 'U',
				Result: i.Wrong,
			},
			i.Clue{
				Letter: 'N',
				Result: i.Wrong,
			},
			i.Clue{
				Letter: 'E',
				Result: i.Wrong,
			},
			i.Clue{
				Letter: 'S',
				Result: i.Correct,
			},
		}
		assert.Equal(t, i.GuessResult{Solved: false, Clues: expectedClues}, res)
	})

	t.Run("first and fifth letters only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("frees")

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
				Letter: 'E',
				Result: i.Wrong,
			},
			i.Clue{
				Letter: 'S',
				Result: i.Correct,
			},
		}
		assert.Equal(t, i.GuessResult{Solved: false, Clues: expectedClues}, res)

	})

	t.Run("all correct", func(t *testing.T) {
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

	})
}
