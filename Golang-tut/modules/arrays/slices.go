package arrays

import "fmt"

func MySlices() {
	cars := make([]string, 3)
	fmt.Println("Car", cars)

	cars[0] = "Bentley"
	cars[1] = "Rolls-Royce"
	cars[2] = "BMW"

	fmt.Println("Set", cars)
	fmt.Println("Get", cars[2])

	fmt.Println("Cars Length:", len(cars))

	cars = append(cars, "Lamborghini")
	cars = append(cars, "Ferrari")
	cars = append(cars, "Skoda Octavia")
	fmt.Println("Cars:", cars)

	lego := cars[2:5]
	fmt.Println("Lego result", lego)

	lego = cars[:5]
	fmt.Println("Lego result 1", lego)

	lego = cars[2:]
	fmt.Println("Lego result 2", lego)

	torque := []string{"G", "H", "I"}
	fmt.Println("Torque:", torque)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerCircle := i + 1
		twoD[i] = make([]int, innerCircle)
		for j := 0; j < innerCircle; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D", twoD)
}
