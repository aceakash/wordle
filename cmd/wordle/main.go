package main

import (
	"flag"
	"fmt"
	"github.com/aceakash/wordle/internal"
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

	game := internal.NewGame(secret)

	var guess string
	_, err := fmt.Fscanln(in, &guess)
	if err != nil {
		return err
	}

	res := game.Guess(guess[:len(guess)-1])
	if res.Solved {
		_, err = fmt.Fprintln(out, "Correct")
		return err
	}
	_, err = fmt.Fprintf(out, "Wrong, the word was %s\n", secret)
	return err
}
