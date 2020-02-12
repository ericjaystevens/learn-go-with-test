package main

import (
	"testing"
)

// Hello says hi
func TestHello(t *testing.T) {
	got := Hello("eric")
	want := "hello, eric"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
