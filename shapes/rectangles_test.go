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
	areaTests := []struct {
		test string
		shape Shape
		want float64
	}{
		{ 
			test : "simple circle area",
			shape: Circle{ radius: 10.0},
			want: 314.1592653589793,
		},
		{
			test : "simple Rectangle area",
			shape : Rectangle{width : 2.0,length : 2.0},
			want : 4.0,
		},
		{
			test : "simple triangle area",
			shape : Triangle{width : 12,height : 6},
			want : 36,
		},
	}

	for _, tt := range areaTests{
		t.Run(tt.test, func(t *testing.T){
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("failed when running %#v got: %g, want %g", tt.shape, got,tt.want)
			}
		})
	}
} 
