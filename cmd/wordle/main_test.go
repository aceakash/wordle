package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWhenFirstGuessIsCorrect(t *testing.T) {
	var stdout bytes.Buffer
	secret := "cater"
	stdin := strings.NewReader(secret)

	err := run(secret, nil, &stdout, stdin)

	assert.Nil(t, err)
	assert.Equal(t, "Correct\n", stdout.String())
}

func TestWhenFirstGuessIsWrong(t *testing.T) {
	var stdout bytes.Buffer
	secret := "cater"
	stdin := strings.NewReader("guess")

	err := run(secret, nil, &stdout, stdin)

	assert.Nil(t, err)
	assert.Equal(t, "Wrong, the word was cater\n", stdout.String())
}
