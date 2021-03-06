package arrays

import (
	"fmt"
)

func MyArrays() {
	var a [5]int
	fmt.Println("a", a)

	a[4] = 100
	fmt.Println("Set", a)
	fmt.Println("Get", a[4])

	fmt.Println("Length", len(a))

	boss := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Boss", boss)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2nd ", twoD)
}
