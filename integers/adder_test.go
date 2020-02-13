package integers

import(
	"testing"
	"fmt"
	"reflect"
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
	numbers := []int{1,2,3}
	got := Sum(numbers)
	want := 6
	if got != want {
		t.Errorf("got: %d, wanted: %d", got, want)
	}
}

func TestSumAll(t *testing.T){
	sliceA := []int{1,2,3}
	sliceB := []int{2,2,1}

	got := SumAll(sliceA, sliceB)
	want := []int{6,5}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got: %d, wanted: %d", got, want)
	}
}