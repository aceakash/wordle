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

		expectedClue := i.Clue{Letter: 'F', Result: i.Correct,}
		assert.Equal(t, expectedClue, res.Clues[0])
	})

	t.Run("third letter only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("ruler")

		expectedClue := i.Clue{Letter: 'L', Result: i.Correct,}
		assert.Equal(t, expectedClue, res.Clues[2])
	})

	t.Run("fifth letter only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("trees")

		expectedClue := i.Clue{Letter: 'S', Result: i.Correct,}
		assert.Equal(t, expectedClue, res.Clues[4])
	})

	t.Run("first and fifth letters only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("frees")

		expectedFirstClue := i.Clue{Letter: 'F', Result: i.Correct,}
		expectedFifthClue := i.Clue{Letter: 'S', Result: i.Correct,}
		assert.Equal(t, expectedFirstClue, res.Clues[0])
		assert.Equal(t, expectedFifthClue, res.Clues[4])
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

func TestGame_MisplacedLettersAreMarked(t *testing.T) {
	
}
