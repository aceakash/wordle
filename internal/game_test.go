package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_ShowSolvedOrNotAfterFirstGuess(t *testing.T) {
	t.Run("when first guess is wrong", func(t *testing.T) {
		g := NewGame("cater")
		res := g.Guess("bravo")

		assert.Equal(t, GuessResult{Solved: false}, res)
	})

	t.Run("when first guess is right", func(t *testing.T) {
		g := NewGame("cater")
		res := g.Guess("cater")

		assert.Equal(t, GuessResult{Solved: true}, res)
	})

}
