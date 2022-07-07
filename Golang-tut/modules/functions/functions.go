package functions

import "fmt"

func Myfunc() {
	result := plus_add(4, 7)
	fmt.Println("Result:", result)
	result = multiply(3, 6, 7)
	fmt.Println("Result:", result)
	sums(45, 67)
	sums(1, 2, 3, 4, 5, 6, 7, 8, 9)
	odd := []int{1, 3, 7, 9}
	sums(odd...)
	square := squareUp()
	fmt.Println(square())
	fmt.Println(square())
	fmt.Println(square())

	squareInts := squareUp()
	fmt.Println(squareInts())
}

func plus_add(a, b int) int {
	return a + b
}

func multiply(a, b, c int) int {
	return a*b*c + 12
}

func sums(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func squareUp() func() int {
	index := 0
	return func() int {
		index++
		return index
	}
}
