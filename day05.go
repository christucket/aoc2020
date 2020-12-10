package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

//   BFFFBBFRRR: row 70, column 7, seat ID 567
// 0b1000110 = 70  0b111 = 7
// convert bs to 1s and f's to 0s, convert Rs to 1s and Ls to 0s

func ConvertToBits(seat string) int {
	var b = 0
	var start = int(math.Pow(2, float64(len(seat))))

	for _, c := range seat {
		start >>= 1
		// B or R
		if c == 66 || c == 82 {
			b |= start
		}
	}

	return b
}

func day05() {
	inp, _ := ioutil.ReadFile("./inputs/day05.input")

	data := GetStringInput(inp)

	max := 0
	all_seats := []int{}
	for _, seat := range data {
		row := ConvertToBits(seat[0:7])
		col := ConvertToBits(seat[7:10])
		seat_id := row*8 + col

		if seat_id > max {
			max = seat_id
		}
		all_seats = append(all_seats, seat_id)
	}
	fmt.Println("max seat id:", max)

	sort.Ints(all_seats)
	i := 0
	for i < len(all_seats)-1 {
		if all_seats[i]+1 != all_seats[i+1] {
			fmt.Println("missing the one in front:", all_seats[i], all_seats[i+1])
		}
		i++
	}

}
