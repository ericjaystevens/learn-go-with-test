package shapes

import(
	"testing"
)

func TestPerRectangle(t *testing.T){
	length := 2
	height := 2
	got := PerRectangle(height,length)
	want := 8

	if got != want {
		t.Errorf("got: %d, want %d", got,want)
	}
}