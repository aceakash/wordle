package main

import (
	"flag"
	"fmt"
	"github.com/aceakash/wordle/internal"
	"github.com/aceakash/wordle/internal/adapters"
	"github.com/fatih/color"
	"io"
	"os"
	"strings"
)

func main() {
	secret := flag.String("secret", "", "Provide a secret word for testing")
	flag.Parse()

	err := run(*secret, nil, os.Stdout, os.Stdin)
	if err != nil {
		panic(err)
	}
}

func run(secret string, otherArgs []string, out io.Writer, in io.Reader) error {
	randomIntPicker := adapters.NewRandomIntPicker()
	secretPicker := adapters.NewFileSecretPicker("wordle-answers-alphabetical.txt", randomIntPicker)
	allowedGuessChecker, err := adapters.NewGuessChecker([]string{"wordle-allowed-guesses.txt", "wordle-answers-alphabetical.txt"})
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "")
	if err != nil {
		panic(err)
	}
	if secret == "" {
		secret, err = secretPicker.PickSecret()
		if err != nil {
			return err
		}
	}
	game, err := internal.NewGame(secret, allowedGuessChecker)
	if err != nil {
		panic(err)
	}

	var guess string

	maxGuesses := 6
	solved := false
	guessesUsed := 0
	for guessesUsed < maxGuesses {
		_, err := fmt.Fprintf(out, "\nGuess (%d of %d) ? ", guessesUsed+1, maxGuesses)
		if err != nil {
			return err
		}
		_, err = fmt.Fscanln(in, &guess)
		if err != nil {
			return err
		}
		res, err := game.Guess(guess)
		if err != nil {
			fmt.Fprintf(out, "\n%s\n", err.Error())
			continue
		}
		err = renderClues(res.Clues, out)
		if err != nil {
			return err
		}
		if res.IsCorrectAnswer() {
			solved = true
			break
		}
		guessesUsed++
	}
	if !solved {
		_, err := fmt.Fprintf(out, "\nThe word was %s\n", strings.ToUpper(secret))
		if err != nil {
			return err
		}
	}
	fmt.Fprintln(out, "")
	return nil
}

func renderClues(clues [5]internal.Clue, stdout io.Writer) error {
	correct := color.New(color.BgGreen, color.FgHiBlack)
	incorrect := color.New(color.BgHiBlack, color.FgWhite)
	misplaced := color.New(color.BgYellow, color.FgHiBlack)
	for i := 0; i < 5; i++ {
		format := " %c "
		if clues[i].Result == internal.Correct {
			_, err := correct.Fprintf(stdout, format, clues[i].Letter)
			if err != nil {
				return err
			}
			continue
		}
		if clues[i].Result == internal.Misplaced {
			_, err := misplaced.Fprintf(stdout, format, clues[i].Letter)
			if err != nil {
				return err
			}
			continue
		}
		_, err := incorrect.Fprintf(stdout, format, clues[i].Letter)
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprintln(stdout, "")
	if err != nil {
		return err
	}
	return nil
}
