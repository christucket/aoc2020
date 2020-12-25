package main

import (
	"fmt"
	"io/ioutil"
)

func TurnToHex(mapping string, x int, y int) (int, int) {
	for pos := 0; pos < len(mapping); pos++ {
		dir := string(mapping[pos])
		if string(mapping[pos]) == "s" || string(mapping[pos]) == "n" {
			pos++
			if string(mapping[pos]) == "e" {
				dir += "e"
			} else {
				dir += "w"
			}
		}

		switch dir {
		case "e":
			x += 1
		case "w":
			x -= 1
		case "se":
			if y%2 != 0 {
				x++
			}
			y++
		case "sw":
			if y%2 == 0 {
				x--
			}
			y++
		case "ne":
			if y%2 != 0 {
				x++
			}
			y--
		case "nw":
			if y%2 == 0 {
				x--
			}
			y--
		default:
			fmt.Println("unknown:", dir)
		}
	}

	return x, y
}

func GetHexNeighbors(grid map[string]int, x int, y int) int {
	dirs := []string{"e", "w", "ne", "nw", "se", "sw"}
	neighbors := 0

	for _, dir := range dirs {
		dx, dy := TurnToHex(dir, x, y)
		key := fmt.Sprintf("%v,%v", dx, dy)
		existing := grid[key]
		if existing == 1 {
			neighbors++
		}
	}

	return neighbors
}
func ConhexsGameOfTiles(grid map[string]int) map[string]int {
	next_day := make(map[string]int)
	for k, v := range grid {
		next_day[k] = v
	}

	for x := -110; x < 110; x++ {
		for y := -110; y < 110; y++ {
			total_neighbors := GetHexNeighbors(grid, x, y)
			key := fmt.Sprintf("%v,%v", x, y)
			v := grid[key]

			if v == 1 && (total_neighbors == 0 || total_neighbors > 2) {
				next_day[key] = 0
			} else if v == 0 && (total_neighbors == 2) {
				next_day[key] = 1
			}
		}
	}

	return next_day
}

func day24() {
	inp, _ := ioutil.ReadFile("./inputs/day24.input")

	data := GetStringInput(inp)

	grid := make(map[string]int) // 0 white, 1 black

	for _, k := range data {
		x, y := TurnToHex(k, 5, 5)
		key := fmt.Sprintf("%v,%v", x, y)
		existing := grid[key]
		if existing == 0 {
			grid[key] = 1
		} else {
			grid[key] = 0
		}
	}

	fmt.Println(grid)

	for times := 0; times < 100; times++ {
		grid = ConhexsGameOfTiles(grid)
		black_count := 0
		for _, tile := range grid {
			if tile == 1 {
				black_count++
			}
		}

		fmt.Printf("day %v: %v\n", times+1, black_count)
	}

}
