package mymaps

import "fmt"

func MyMaps() {
	users := make(map[string]int)

	users["kyle"] = 15
	users["Allison"] = 18
	users["Justin"] = 23
	users["Omari"] = 16

	// fmt.Println("Users", users)
	fmt.Println("Users length:", len(users))

	user := users["Victor"]
	fmt.Println("User:", user)

	fmt.Println("Length:", len(users))

	delete(users, "Kyle")
	fmt.Println("Users:", users)

	_, news := users["Kyle"]
	fmt.Println("News:", news)

	africans := map[string]int{
		"Nigeria":      50,
		"Ghana":        45,
		"Liberia":      30,
		"Sierra Leone": 25,
	}

	for countries, index := range africans {
		fmt.Println("African Countries:", countries, "->", "Index", index)
	}
}
