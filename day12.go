package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Movement struct {
	Direction string
	Amount    int
}

func day12() {
	inp, _ := ioutil.ReadFile("./inputs/day12.input")

	data := GetStringInput(inp)

	moves := []Movement{}
	for _, mv := range data {
		amt, _ := strconv.Atoi(mv[1:])
		moves = append(moves, Movement{string(mv[0]), amt})
	}

	// N - 0
	// E - 1
	// S - 2
	// W - 3

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	facing := 1
	x, y := 0, 0
	for _, mv := range moves {
		var go_direction []int
		if mv.Direction == "N" {
			go_direction = directions[0]
		} else if mv.Direction == "E" {
			go_direction = directions[1]
		} else if mv.Direction == "S" {
			go_direction = directions[2]
		} else if mv.Direction == "W" {
			go_direction = directions[3]
		} else if mv.Direction == "L" {
			facing = (facing - mv.Amount/90 + 4) % 4
			go_direction = directions[facing]
		} else if mv.Direction == "R" {
			facing = (facing + mv.Amount/90 + 4) % 4
			go_direction = directions[facing]
		} else if mv.Direction == "F" {
			go_direction = directions[facing]
		}

		if mv.Direction != "R" && mv.Direction != "L" {
			x += go_direction[0] * mv.Amount
			y += go_direction[1] * mv.Amount
		}
	}

	fmt.Println("part 1 x, y:", math.Abs(float64(x)), math.Abs(float64(y)), math.Abs(float64(x))+math.Abs(float64(y)))

	facing = 1
	x, y = 0, 0
	wx, wy := 10, 1
	for _, mv := range moves {
		var go_direction []int
		if mv.Direction == "N" {
			go_direction = directions[0]
		} else if mv.Direction == "E" {
			go_direction = directions[1]
		} else if mv.Direction == "S" {
			go_direction = directions[2]
		} else if mv.Direction == "W" {
			go_direction = directions[3]
		} else if strings.Contains("LR", mv.Direction) {
			times := mv.Amount / 90
			if mv.Direction == "L" {
				times *= 3 // three rights is a left, right??
			}
			for times > 0 {
				wx, wy = wy, -wx
				times--
			}
		} else if mv.Direction == "F" {
			go_direction = directions[facing]
		}

		if strings.Contains("NESW", mv.Direction) {
			wx += go_direction[0] * mv.Amount
			wy += go_direction[1] * mv.Amount
		}
		if mv.Direction == "F" {
			x += wx * mv.Amount
			y += wy * mv.Amount
		}
	}
	fmt.Println("part 2 x, y:", math.Abs(float64(x)), math.Abs(float64(y)), math.Abs(float64(x))+math.Abs(float64(y)))

}
