package myrunes

import (
	"fmt"
	"unicode/utf8"
)

func MyRunes() {
	const s = "hello world"

	fmt.Println("Length:", len(s))
	fmt.Println("RuneCounts:", utf8.RuneCountInString(s))
}
