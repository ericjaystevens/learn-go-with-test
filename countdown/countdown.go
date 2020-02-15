package main

import (
	"time"
	"os"
	"io"
	"fmt"
)

type Sleeper interface{
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

const countdownStart = 3
const actionWord = "Go!"

func main()  {
	d := DefaultSleeper{}
	Countdown(os.Stdout, &d)
}

func Countdown(writer io.Writer, sleep Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleep.Sleep()
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, actionWord)
}
