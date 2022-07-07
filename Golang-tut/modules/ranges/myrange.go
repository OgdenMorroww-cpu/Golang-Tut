package ranges

import "fmt"

func MyRanges() {
	even := []int{2, 4, 6}
	sum := 0
	for _, num := range even {
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range even {
		if num == 3 {
			fmt.Println("Index:", i)
		}
	}

	fruits := map[string]string{"a": "apple", "b": "banana", "o": "orange"}
	for keys, values := range fruits {
		fmt.Printf("%s -> %s\n", keys, values)
	}

	for keys := range fruits {
		fmt.Println("Key:", keys)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
