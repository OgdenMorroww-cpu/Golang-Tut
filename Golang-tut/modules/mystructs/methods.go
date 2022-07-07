package mystructs

import (
	"fmt"
)

type Rect struct {
	width  int
	height int
}

func (r *Rect) area() int {
	return r.width * r.height
}

func (r Rect) perim() int {
	return 2*r.width + 2*r.height
}

func MyMethods() {
	r := Rect{width: 20, height: 15}
	fmt.Println("Area:", r.area())
	fmt.Println("Perim:", r.perim())

	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Perim:", rp.perim())
}
