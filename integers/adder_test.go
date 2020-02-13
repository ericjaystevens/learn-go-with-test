package integers

import(
	"testing"
	"fmt"
)

func TestAdder(t *testing.T){
	got := Adder(2,2)
	want := 4

	if got != want {
		t.Errorf("got: %d, want %d", got, want)
	}
}

func ExampleAdder(){
	got := Adder(2,3)
	fmt.Println(got)
	// output: 5
}