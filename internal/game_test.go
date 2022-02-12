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

		expectedClue := i.Clue{Letter: 'F', Result: i.Correct}
		assert.Equal(t, expectedClue, res.Clues[0])
	})

	t.Run("third letter only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("ruler")

		expectedClue := i.Clue{Letter: 'L', Result: i.Correct}
		assert.Equal(t, expectedClue, res.Clues[2])
	})

	t.Run("fifth letter only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("trees")

		expectedClue := i.Clue{Letter: 'S', Result: i.Correct}
		assert.Equal(t, expectedClue, res.Clues[4])
	})

	t.Run("first and fifth letters only", func(t *testing.T) {
		g, _ := i.NewGame("folds")
		res := g.Guess("frees")

		expectedFirstClue := i.Clue{Letter: 'F', Result: i.Correct}
		expectedFifthClue := i.Clue{Letter: 'S', Result: i.Correct}
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
	t.Run("first letter only", func(t *testing.T) {
		g, _ := i.NewGame("wetly")
		res := g.Guess("train")

		expectedClue := i.Clue{Letter: 'T', Result: i.Misplaced}
		assert.Equal(t, expectedClue, res.Clues[0])
	})

	t.Run("fifth letter only", func(t *testing.T) {
		g, _ := i.NewGame("wetly")
		res := g.Guess("yanks")

		expectedClue := i.Clue{Letter: 'Y', Result: i.Misplaced}
		assert.Equal(t, expectedClue, res.Clues[0])
	})

	t.Run("all letters", func(t *testing.T) {
		g, _ := i.NewGame("cater")
		res := g.Guess("trace")

		assert.Equal(t, i.Clue{Letter: 'T', Result: i.Misplaced}, res.Clues[0])
		assert.Equal(t, i.Clue{Letter: 'R', Result: i.Misplaced}, res.Clues[1])
		assert.Equal(t, i.Clue{Letter: 'A', Result: i.Misplaced}, res.Clues[2])
		assert.Equal(t, i.Clue{Letter: 'C', Result: i.Misplaced}, res.Clues[3])
		assert.Equal(t, i.Clue{Letter: 'E', Result: i.Misplaced}, res.Clues[4])
	})
}

func TestGame_WrongLettersAreMarked(t *testing.T) {
	t.Run("first letter only", func(t *testing.T) {
		g, _ := i.NewGame("weary")
		res := g.Guess("train")

		expectedClue := i.Clue{Letter: 'T', Result: i.Wrong}
		assert.Equal(t, expectedClue, res.Clues[0])
	})

	t.Run("fifth letter only", func(t *testing.T) {
		g, _ := i.NewGame("wetly")
		res := g.Guess("float")

		expectedClue := i.Clue{Letter: 'T', Result: i.Misplaced}
		assert.Equal(t, expectedClue, res.Clues[4])
	})

	t.Run("all letters", func(t *testing.T) {
		g, _ := i.NewGame("cater")
		res := g.Guess("lions")

		assert.Equal(t, i.Clue{Letter: 'L', Result: i.Wrong}, res.Clues[0])
		assert.Equal(t, i.Clue{Letter: 'I', Result: i.Wrong}, res.Clues[1])
		assert.Equal(t, i.Clue{Letter: 'O', Result: i.Wrong}, res.Clues[2])
		assert.Equal(t, i.Clue{Letter: 'N', Result: i.Wrong}, res.Clues[3])
		assert.Equal(t, i.Clue{Letter: 'S', Result: i.Wrong}, res.Clues[4])
	})
}
