package myinterfaces

import (
	"fmt"
)

type Base struct {
	num int
}

func (c Base) describe() string {
	return fmt.Sprintf("Base with num=%v", c.num)
}

type Container struct {
	Base
	str string
}

type Continent struct {
	continet_name   string
	region          string
	total_countries int
	Nigeria
}

type Nigeria struct {
	name         string
	president    string
	independence int
}

func (con Continent) details(colour_race string) string {
	return colour_race
}

func MyStructEm() {
	co := Container{
		Base: Base{
			num: 1,
		},
		str: "Some girl",
	}
	fmt.Printf("Co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("Also num:", co.Base.num)
	fmt.Println("Describe:", co.describe())

	type Describer interface {
		describe() string
	}
	var d Describer = co
	fmt.Println("Describer:", d.describe())
	continents := Continent{
		continet_name:   "Africa",
		region:          "West Africa",
		total_countries: 58,
		Nigeria: Nigeria{
			name:         "Nigeria",
			president:    "Buhari",
			independence: 1960,
		},
	}
	fmt.Println("Race:", continents.details("Negroid Human Race"))
	fmt.Println("Continent Name:", continents.continet_name)
	fmt.Println("Region:", continents.region)
	fmt.Println("Number of countries:", continents.total_countries)
	fmt.Println("Country Name:", continents.Nigeria.name)
	fmt.Println("Country President:", continents.Nigeria.president)
	fmt.Println("Independence Date:", continents.independence)

}
