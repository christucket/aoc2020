package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetWorld(world map[string]bool, x int, y int, w int, z int) bool {
	val := world[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)]

	return val
}

func SetWorld(world map[string]bool, x int, y int, z int, w int, val bool) {
	world[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)] = val
}

func PrintWorld(world map[string]bool) {
	for z := -10; z < 10; z++ {
		world_string := fmt.Sprintf("z=%d\n", z)
		for y := -10; y < 10; y++ {
			row := ""
			for x := -10; x < 10; x++ {
				on := GetWorld(world, x, y, z, 0)
				char := ""
				if on {
					char = "#"
				} else if x == 0 || y == 0 {
					char = "-"
				}
				row += char
			}
			if strings.Contains(row, "#") {
				world_string += row + "\n"
			}
		}
		if strings.Contains(world_string, "#") {
			fmt.Println(world_string)
		}
	}
	fmt.Println("==================")
}

func GetWorldNeighbors(world map[string]bool, x int, y int, z int, w int) int {
	neighbors := 0

	for dw := w - 1; dw <= w+1; dw++ {
		for dz := z - 1; dz <= z+1; dz++ {
			for dy := y - 1; dy <= y+1; dy++ {
				for dx := x - 1; dx <= x+1; dx++ {
					if GetWorld(world, dx, dy, dz, dw) &&
						!(x == dx && y == dy && z == dz && w == dw) {
						neighbors++
					}
				}
			}
		}
	}

	return neighbors
}

func day17() {
	debug := false
	inp, _ := ioutil.ReadFile("./inputs/day17.input")

	data := GetStringInput(inp)

	world := make(map[string]bool)

	for y, line := range data {
		for x, cube := range line {
			SetWorld(world, x, y, 0, 0, string(cube) == "#")
		}
	}

	times := 0
	size := 9 // could set a specific size for minX, maxX, minY, maxY, etc
	// and calculate it in the loop when setting an active cube
	for times < 6 {
		if debug {
			PrintWorld(world)
		}
		world_backup := make(map[string]bool)
		for w := -size + 6; w < size-6; w++ {
			for z := -size + 6; z < size-6; z++ {
				for y := -size; y < size; y++ {
					for x := -size; x < size; x++ {
						neighbors := GetWorldNeighbors(world, x, y, z, w)
						position := GetWorld(world, x, y, z, w)
						if position {
							if neighbors == 2 || neighbors == 3 {
								SetWorld(world_backup, x, y, z, w, true)
							}
						} else {
							if neighbors == 3 {
								SetWorld(world_backup, x, y, z, w, true)
							}
						}
					}
				}
			}
		}
		size++
		times++
		world = world_backup

		count := 0
		for _, v := range world {
			if v {
				count++
			}
		}
		fmt.Println("current active:", count)
	}
}
