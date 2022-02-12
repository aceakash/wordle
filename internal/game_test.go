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

func TestGame_ShowSolvedOrNotAfterFirstGuess(t *testing.T) {
	t.Run("when first guess is wrong", func(t *testing.T) {
		g, _ := i.NewGame("cater")
		res := g.Guess("bravo")

		assert.Equal(t, i.GuessResult{Solved: false}, res)
	})

	t.Run("when first guess is right", func(t *testing.T) {
		g, _ := i.NewGame("cater")
		res := g.Guess("cater")

		assert.Equal(t, i.GuessResult{Solved: true}, res)
	})

}

func TestGame_WhenOnlyFirstLetterIsCorrect(t *testing.T) {
	g, _ := i.NewGame("folds")
	res := g.Guess("freak")

	expectedFirstClue := i.Clue{
		Letter: "F",
		Result: i.Correct,
	}
	expectedClues := [5]i.Clue{expectedFirstClue}
	assert.Equal(t, i.GuessResult{Solved: false, Clues: expectedClues}, res)
}
