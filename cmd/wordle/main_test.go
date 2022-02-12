package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecretWordIsPrinted(t *testing.T) {
	var stdout bytes.Buffer

	err := run("akash", nil, &stdout, nil)

	assert.Nil(t, err)
	assert.Equal(t, "akash\n", stdout.String())
}
