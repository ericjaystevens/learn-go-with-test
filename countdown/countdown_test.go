package main

import(
	"testing"
	"bytes"
)	

type spySleeper struct{
	Calls int
}

func (s spySleeper) Sleep(){
	s.Calls++
}

func TestCountdown(t *testing.T)  {
	buffer := bytes.Buffer{}
	sleeper := spySleeper{}
	Countdown(&buffer,&sleeper)
	got := buffer.String()
    want := `3
2
1
Go!`

	if got != want{
		t.Errorf("got %s wanted %s",got,want)
	}
}