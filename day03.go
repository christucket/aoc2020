package main

import (
	"fmt"
	"io/ioutil"
)

func CheckTreesWithSlope(data []string, dx int, dy int) int {
	x, y := 0, 0
	trees := 0

	for y+dy < len(data) {
		x = (x + dx) % len(data[0])
		y += dy

		spot := string(data[y][x])

		if spot == "#" {
			trees++
		}
	}

	return trees
}

// Starting at the top-left corner of your map
// and following a slope of right 3 and down 1,
// how many trees would you encounter
func day03() {
	inp, _ := ioutil.ReadFile("./day03.input")

	data := GetStringInput(inp)

	slope11 := CheckTreesWithSlope(data, 1, 1)
	slope31 := CheckTreesWithSlope(data, 3, 1)
	slope51 := CheckTreesWithSlope(data, 5, 1)
	slope71 := CheckTreesWithSlope(data, 7, 1)
	slope12 := CheckTreesWithSlope(data, 1, 2)

	fmt.Println("1x1", slope11)
	fmt.Println("3x1", slope31)
	fmt.Println("5x1", slope51)
	fmt.Println("7x1", slope71)
	fmt.Println("1x2", slope12)

	fmt.Println("together", slope11*slope12*slope31*slope51*slope71)
}
