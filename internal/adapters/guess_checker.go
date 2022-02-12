package adapters

import "strings"

type GuessChecker struct {
	words map[string]bool
}

func (g GuessChecker) IsGuessAllowed(guess string) bool {
	_, found := g.words[guess]
	return found
}

func NewGuessChecker(filePaths []string) (GuessChecker, error) {
	wordMap := map[string]bool{}
	for _, filePath := range filePaths {
		words, err := ReadWordsFromFile(filePath)
		if err != nil {
			return GuessChecker{}, err
		}
		for _, word := range words {
			wordMap[strings.ToUpper(word)] = true
		}
	}

	return GuessChecker{wordMap}, nil
}
