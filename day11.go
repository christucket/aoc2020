package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CountNearbySeated(data [][]string, x int, y int) int {
	checks := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*, {0, 0}*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	seated := 0

	for _, xy := range checks {
		dx, dy := xy[0], xy[1]

		loop_again := true
		multiplier := 1 // loop and multipler are for part 2 only

		for loop_again {
			if x+dx*multiplier >= 0 && x+dx*multiplier < len(data[0]) {
				if y+dy*multiplier >= 0 && y+dy*multiplier < len(data) {
					if data[y+dy*multiplier][x+dx*multiplier] == "#" {
						seated++
						loop_again = false
					} else if data[y+dy*multiplier][x+dx*multiplier] == "L" {
						loop_again = false
					}
				} else {
					loop_again = false
				}
			} else {
				loop_again = false
			}
			multiplier++
		}
	}
	return seated
}

func ComputeLayoutDifference(data [][]string) (bool, [][]string) {
	var seats [][]string
	changed := false

	for y, row := range data {
		var new_row []string

		for x, _ := range row {
			seated_nearby := CountNearbySeated(data, x, y)
			if data[y][x] == "L" && seated_nearby == 0 {
				new_row = append(new_row, "#")
				changed = true
			} else if data[y][x] == "#" && seated_nearby >= 5 { // change to 4 for part 1
				new_row = append(new_row, "L")
				changed = true
			} else {
				new_row = append(new_row, data[y][x])
			}
		}

		seats = append(seats, new_row)
	}

	return changed, seats
}

func day11() {
	inp, _ := ioutil.ReadFile("./inputs/day11.input")

	data := Get2DStringInput(inp)

	changed := true
	time := 0
	for changed {
		changed, data = ComputeLayoutDifference(data)

		// just in case
		time++
		if time > 100 {
			return
		}
	}

	all_seats := ""
	for _, row := range data {
		all_seats = all_seats + " " + strings.Join(row, "")
	}
	fmt.Println(strings.Count(all_seats, "#"), all_seats)
}
