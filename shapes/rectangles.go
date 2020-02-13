package shapes

import(
	"math"
)
type Rectangle struct {
	length float64
	width float64
}

type Circle struct {
	radius float64
}

// PerRectangle returnts the perimeter of a rectangel based on length and height
func (r Rectangle )PerRectangle() (perimter float64){
	perimter = 2 * r.width + 2 * r.length
	return
}

func (c Circle) Area() (area float64){
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() (area float64){
	return (r.length + r.width)
}