package main

import (
	"flag"
	"fmt"
	"github.com/aceakash/wordle/internal"
	"github.com/aceakash/wordle/internal/adapters"
	"github.com/fatih/color"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	secret := flag.String("secret", "", "Provide a secret word for testing")
	flag.Parse()
	randomIntPicker := NewRandomIntPicker()
	secretPicker := adapters.NewFileSecretPicker("wordle-answers-alphabetical.txt", randomIntPicker)

	err := run(*secret, secretPicker, nil, os.Stdout, os.Stdin)
	if err != nil {
		panic(err)
	}
}

func NewRandomIntPicker() RandomIntPicker {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return RandomIntPicker{r}
}

type RandomIntPicker struct {
	rand *rand.Rand
}

func (r RandomIntPicker) PickRandomInt(upto int) int {
	return r.rand.Intn(upto)
}

type SecretPicker interface {
	PickSecret() (string, error)
}

func run(secret string, secretPicker SecretPicker, otherArgs []string, out io.Writer, in io.Reader) error {
	_, err := fmt.Fprintf(out, "")
	if err != nil {
		panic(err)
	}
	if secret == "" {
		secret, err = secretPicker.PickSecret()
		if err != nil {
			return err
		}
	}
	game, err := internal.NewGame(secret)
	if err != nil {
		panic(err)
	}

	var guess string

	maxGuesses := 6
	solved := false
	for i := 0; i < maxGuesses; i++ {
		_, err := fmt.Fprintf(out, "\nGuess (%d of %d) ? ", i+1, maxGuesses)
		if err != nil {
			return err
		}
		_, err = fmt.Fscanln(in, &guess)
		if err != nil {
			return err
		}
		res := game.Guess(guess)
		err = renderClues(res.Clues, out)
		if err != nil {
			return err
		}
		if res.IsCorrectAnswer() {
			solved = true
			break
		}
	}
	if !solved {
		_, err := fmt.Fprintf(out, "\nThe word was %s\n", strings.ToUpper(secret))
		if err != nil {
			return err
		}
	}
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
