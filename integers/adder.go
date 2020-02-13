package integers

import(
)

// Adder adds some thing together
func Adder(x int, y int) int{
	sum := x + y
	return sum
}

func Sum(numbers [3]int) (sum int){
	
	for i := range numbers {
		sum += numbers[i]	
	}

	return
}