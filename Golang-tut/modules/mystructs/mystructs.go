package mystructs

import (
	"fmt"
)

type Country struct {
	name         string
	independence int
	isDemocratic bool
	president    string
}

func country(name string) *Country {
	c := Country{name: name}
	c.independence = 1963
	return &c
}

func MyStructs() {
	fmt.Println(Country{"Nigeria", 1960, true, "Buhari"})
	fmt.Println(Country{"Ghana", 1958, true, "Nana Akufudo"})
	fmt.Println(Country{"United States", 1776, true, "Joe Biden"})
	fmt.Println(Country{"China", 1945, false, "XI Jin Ping"})
	fmt.Println(country("Austria"))

	g := Country{name: "India", independence: 1943, isDemocratic: true, president: "Modi"}
	fmt.Println("Country Name:", g.name)
	fmt.Println("Independence Date:", g.independence)
	fmt.Println("Democratic:", g.isDemocratic)
	fmt.Println("President:", g.president)
}
