package conditionals

import (
	"fmt"
	"time"
)

func SwitchStatement() {
	i := 2
	fmt.Println("Write ", i, "  as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatTheHeck := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am an int")
		default:
			fmt.Println("Don't know the types", t)
		}
	}
	whatTheHeck(true)
	whatTheHeck(2)
	whatTheHeck("hey")
}
