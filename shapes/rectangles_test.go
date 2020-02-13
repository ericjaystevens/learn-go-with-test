package shapes

import(
	"testing"
)


func TestPerRectangle(t *testing.T){
	rec := Rectangle{2,2}
	got := rec.PerRectangle()
	want := 8.0

	if got != want {
		t.Errorf("got: %.2f, want %.2f", got,want)
	}
}

func TestArea(t *testing.T){
	t.Run("rectange", func(t *testing.T){
		rec := Rectangle{2.0,2.0}
		got := rec.Area()
		want := 4.0

		if got != want {
			t.Errorf("got: %.2f, want %.2f", got,want)
		}
	})
	t.Run("circle", func(t *testing.T){
		cir := Circle{10.0}
		got := cir.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got: %g, want %g", got,want)
		}
	})
} 
