package shapes

import(
	"testing"
)

type Shape interface{
	Area() float64
}

func TestPerRectangle(t *testing.T){
	rec := Rectangle{2,2}
	got := rec.PerRectangle()
	want := 8.0

	if got != want {
		t.Errorf("got: %.2f, want %.2f", got,want)
	}
}

func TestArea(t *testing.T){
	getArea := func(t *testing.T, shape Shape, want float64){
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got: %g, want %g", got,want)
		}
	}
	t.Run("rectange", func(t *testing.T){
		rec := Rectangle{2.0,2.0}
		getArea(t,rec,4.0)
	})
	t.Run("circle", func(t *testing.T){
		cir := Circle{10.0}
		getArea(t,cir,314.1592653589793)
	})
} 
