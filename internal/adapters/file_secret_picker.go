package adapters

import (
	"os"
	"strings"
)

type IRandomIntPicker interface {
	PickRandomInt(upto int) int
}

type FileSecretPicker struct {
	filePath        string
	randomIntPicker IRandomIntPicker
}

func (f FileSecretPicker) PickSecret() (string, error) {
	bytes, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", nil
	}
	words := strings.Split(string(bytes), "\n")
	index := f.randomIntPicker.PickRandomInt(len(words))
	return words[index], nil
}

func NewFileSecretPicker(filePath string, randomIntPicker IRandomIntPicker) FileSecretPicker {
	return FileSecretPicker{filePath, randomIntPicker}
}
