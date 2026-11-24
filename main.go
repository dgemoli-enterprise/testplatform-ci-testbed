package main

import (
	"fmt"
	"io"
	"os"
)

func greet(w io.Writer) error {
	_, err := fmt.Fprintln(w, "hello world")
	return err
}

func main() {
	if err := greet(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
