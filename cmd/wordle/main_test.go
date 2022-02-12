package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type DummySecretPicker struct{}

func (d DummySecretPicker) PickSecret() (string, error) {
	panic("implement me")
}

func TestWhenFirstGuessIsCorrect(t *testing.T) {
	t.SkipNow()
	var stdout bytes.Buffer
	secret := "cater"
	stdin := strings.NewReader(secret)

	err := run(secret, DummySecretPicker{}, nil, &stdout, stdin)

	assert.Nil(t, err)
	assert.Equal(t, "Correct\n", stdout.String())
}

func TestWhenFirstGuessIsWrong(t *testing.T) {
	t.SkipNow()
	var stdout bytes.Buffer
	secret := "cater"
	stdin := strings.NewReader("guess")

	err := run(secret, DummySecretPicker{}, nil, &stdout, stdin)

	assert.Nil(t, err)
	assert.Equal(t, "Wrong, the word was cater\n", stdout.String())
}
