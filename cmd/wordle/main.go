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
	_, err := fmt.Fprintln(out, secret)
	return err
}
