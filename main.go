package main

import "fmt"

type figure struct {
	height, width int
	name          string
}

func main() {
	triangle := figure{5, 10, "Triangle"}
	area(triangle)

	square := figure{6, 6, "Square"}
	area(square)

	rectangle := figure{5, 10, "Rectangle"}
	area(rectangle)
}

func area(shape figure) {
	var area int
	if shape.name == "Triangle" {
		area = shape.height * shape.width / 2
	} else {
		area = shape.height * shape.width
	}

	fmt.Println(shape.name, area, "area")
}
