package integers

import(
)

// Adder adds some thing together
func Adder(x int, y int) int{
	sum := x + y
	return sum
}

func Sum(numbers []int) (sum int){
	
	for i := range numbers {
		sum += numbers[i]	
	}

	return
}

func SumAll(numbers ...[]int) []int{
	var sums []int
	for _, number := range numbers{
		sums = append(sums, Sum(number))
	}
	return sums
}