package main

import (
	"flag"
	"fmt"
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
	var guess string
	_, err := fmt.Fscanln(in, &guess)
	if err != nil {
		return err
	}
	if guess == secret {
		_, err = fmt.Fprintln(out, "Correct")
		return err
	}
	_, err = fmt.Fprintf(out, "Wrong, the word was %s\n", secret)
	return err
}
