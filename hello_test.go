package main

import (
	"testing"
)

// Hello says hi
func TestHello(t *testing.T) {
	assertCorrect := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, when we really wanted %q", got, want)
		}
	}
	t.Run("Say hello based on name", func(t *testing.T) {
		got := Hello("eric")
		want := "hello, eric"
		assertCorrect(t, got, want)
	})
	t.Run("Say hello without on name", func(t *testing.T) {
		got := Hello("")
		want := "hi, you!"
		assertCorrect(t, got, want)

	})

}
