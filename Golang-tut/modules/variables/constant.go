package variables

import (
	"fmt"
	"math"
)

const s string = "constant"

func Constants() {
	fmt.Println(s)

	const n = 50000_0000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
