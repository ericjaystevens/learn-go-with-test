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

func TestAreaRectangle(t *testing.T){
	rec := Rectangle{2.0,2.0}
	got := AreaRectangle(rec)
	want := 4.0

	if got != want {
		t.Errorf("got: %.2f, want %.2f", got,want)
	}
} 