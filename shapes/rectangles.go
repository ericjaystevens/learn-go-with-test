package shapes

import(
)
type Rectangle struct {
	length float64
	width float64
}

// PerRectangle returnts the perimeter of a rectangel based on length and height
func PerRectangle(height int, length int ) (perimter int){
	perimter = 2 * height + 2 * length
	return
}

func AreaRectangle(rec Rectangle) (area float64){
	return rec.width * rec.length
}