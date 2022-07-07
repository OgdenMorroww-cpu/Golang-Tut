package pointers

import "fmt"

func MyPointers() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("ZeroVal:", i)

	zerovalPtr(&i)
	fmt.Println("ZeroPtr:", i)

	fmt.Println("Pointer:", &i)
}

func zeroval(ival int) {
	ival = 0
}

func zerovalPtr(iptr *int) {
	*iptr = 0
}
