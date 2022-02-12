package adapters

import (
	"os"
	"strings"
)

func ReadWordsFromFile(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), "\n"), nil
}
