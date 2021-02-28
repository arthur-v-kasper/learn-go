package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	//The buffer type from the bytes package implements the Writer interface.
	buffer := bytes.Buffer{}

	Greet(&buffer, "Kasper")

	got := buffer.String()
	want := "Hello, Kasper"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
