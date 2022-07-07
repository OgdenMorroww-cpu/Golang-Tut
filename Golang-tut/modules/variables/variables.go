package variables

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c, d int = 4, 6, 8
	fmt.Println(b, c, d)

	var truth = true
	fmt.Println(truth)

	var exponetial int
	fmt.Println(exponetial)

	company := "Apple Corp"
	fmt.Println("The company name is: ", company)
}
