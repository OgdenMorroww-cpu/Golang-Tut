package variables

import (
	"fmt"
	"strings"
)

func Calculator() {
	var inputs = "Num1"
	fmt.Println(inputs)
	var num1 int
	fmt.Scanln(&num1)

	var inputs2 = "Num2"
	fmt.Println(inputs2)
	var num2 int
	fmt.Scanln(&num2)

	var calculator_operator = "calculator operator"
	fmt.Println(calculator_operator)
	var commands string
	fmt.Scanln(&commands)

	if strings.ToUpper(commands) == "+" {
		addition := num1 + num2
		fmt.Println("Addition:", addition)
	} else if strings.ToUpper(commands) == "-" {
		subtraction := num1 - num2
		fmt.Println("Subtraction:", subtraction)
	} else if strings.ToUpper(commands) == "*" {
		multiplication := num1 * num2
		fmt.Println("Multiplication:", multiplication)
	} else if strings.ToUpper(commands) == "/" {
		division := num1 / num2
		fmt.Println("Division:", division)
	} else if strings.ToUpper(commands) == "%" {
		remainder := num1 % num2
		fmt.Println("Remainder:", remainder)
	} else {
		fmt.Println("Invalid Commands")
	}
}
