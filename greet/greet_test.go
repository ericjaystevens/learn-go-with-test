package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dug")
	got := buffer.String()
	want := "hi ya Dug"

	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}
