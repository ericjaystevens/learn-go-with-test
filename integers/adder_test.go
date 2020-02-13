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

func TestSum(t *testing.T){
	numbers := [3]int{1,2,3}
	got := Sum(numbers)
	want := 6
	if got != want {
		t.Errorf("got: %d, wanted: %d", got, want)
	}
}