package greet

import (
	"bytes"
	"fmt"
)

func Greet(output *bytes.Buffer, name string) {
	fmt.Fprintf(output, "hi ya %s", name)
}
