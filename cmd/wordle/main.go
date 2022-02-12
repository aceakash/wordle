package main

import (
	"flag"
	"fmt"
	"github.com/aceakash/wordle/internal"
	"github.com/fatih/color"
	"io"
	"os"
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

	game, err := internal.NewGame(secret)
	if err != nil {
		panic(err)
	}

	var guess string
	_, err = fmt.Fscanln(in, &guess)
	if err != nil {
		return err
	}

	res := game.Guess(guess)
	return renderClues(res.Clues, out)
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
