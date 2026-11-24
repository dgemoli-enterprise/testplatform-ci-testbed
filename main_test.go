package main

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	var out strings.Builder
	if err := greet(&out); err != nil {
		t.Fatal(err)
	}
	if got, want := out.String(), "hello world\n"; got != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}

func TestGreetingWriteError(t *testing.T) {
	r, w := io.Pipe()
	r.Close()
	defer w.Close()
	err := greet(w)
	if !errors.Is(err, io.ErrClosedPipe) {
		t.Errorf("error = %v, want %v", err, io.ErrClosedPipe)
	}
}
