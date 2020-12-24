package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func FindCurrentCupIndex(current_cup int, ring []int) int {
	for i := range ring {
		if ring[i] == current_cup {
			return i
		}
	}
	return -1
}

func GetNextThreeCups(current_cup_index int, ring []int) []int {
	next_three_cups := []int{}
	for _, k := range []int{1, 2, 3} {
		idx := (current_cup_index + k) % len(ring)
		next_three_cups = append(next_three_cups, ring[idx])
		ring[idx] = -1
	}

	return next_three_cups
}

func FindDestinationCup(finding_cup int, ring []int, next_three_cups []int) (int, int) {
	i := 0
	if finding_cup == 0 {
		finding_cup = len(ring)
	}

	f := true
	for f {
		f = false
		for _, k := range next_three_cups {
			if finding_cup == k {
				finding_cup--
				if finding_cup == 0 {
					finding_cup = len(ring)
				}
				f = true
			}
		}
	}
	for i < len(ring) {
		// fmt.Println("trying to find", finding_cup)
		if ring[i] == finding_cup {
			return i, finding_cup
		}
		i++
	}

	fmt.Println("couldnt find ", finding_cup, next_three_cups)

	return -1, -1
}

func PlayCupsGame(ring []int, number_of_moves int) []int {
	current_cup := ring[0]

	move := 1
	for move <= number_of_moves {
		current_cup_index := FindCurrentCupIndex(current_cup, ring)
		// fmt.Println(" -> move [", move, "]   ")
		// fmt.Println("cups: ", ring, current_cup, current_cup_index)
		next_three_cups := GetNextThreeCups(current_cup_index, ring)
		_, destination_cup := FindDestinationCup(current_cup-1, ring, next_three_cups)

		if destination_cup == 1 {
			fmt.Println("something after 1", next_three_cups)
		}
		// fmt.Println("pick up: ", next_three_cups)
		// fmt.Println("destnation: ", destination_cup, destination_cup_index)
		// fmt.Println("")
		next_ring := []int{}
		for _, k := range ring {
			if k != -1 {
				next_ring = append(next_ring, k)
			}
			if k == destination_cup {
				next_ring = append(next_ring, next_three_cups...)
			}
		}

		ring = next_ring
		current_cup_index = FindCurrentCupIndex(current_cup, ring)
		current_cup = ring[(current_cup_index+1)%len(ring)]

		move++
		if move%5 == 0 {
			fmt.Println("on move:", move)
		}
	}

	return ring
}

func day23() {
	inp, _ := ioutil.ReadFile("./inputs/day23.sample")

	data := GetStringInput(inp)
	ring := []int{}

	for _, k := range data[0] {
		next_int, _ := strconv.Atoi(string(k))
		ring = append(ring, next_int)
	}

	final_ring := PlayCupsGame(append([]int{}, ring...), 100)

	final_string := ""
	one_index := FindCurrentCupIndex(1, final_ring)
	for len(final_string) < 8 {
		number := strconv.Itoa(final_ring[(one_index+1+len(final_string))%9])
		final_string += number
	}

	fmt.Println("final string:", final_string)
	i := 10
	for i <= 1_000_000 {
		ring = append(ring, i)
		i++
	}
	fmt.Println("starting game")
	final_ring = PlayCupsGame(append([]int{}, ring...), 10_000_000)
	fmt.Println("starting game")

	one_index = FindCurrentCupIndex(1, final_ring)
	fmt.Println(final_ring[one_index+1], final_ring[one_index+2])
}
