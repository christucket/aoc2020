package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

func GetAllCombos(data []int, start_idx int, cache map[int]int) int {
	c := 0
	if start_idx == len(data)-1 {
		return 1
	}

	if start_idx+1 < len(data) && data[start_idx+1]-data[start_idx] < 4 {
		if val, found := cache[start_idx+1]; found {
			c += val
		} else {
			result := GetAllCombos(data, start_idx+1, cache)
			cache[start_idx+1] = result
			c += result
		}
	}
	if start_idx+2 < len(data) && data[start_idx+2]-data[start_idx] < 4 {
		if val, found := cache[start_idx+2]; found {
			c += val
		} else {
			result := GetAllCombos(data, start_idx+2, cache)
			cache[start_idx+2] = result
			c += result
		}
	}
	if start_idx+3 < len(data) && data[start_idx+3]-data[start_idx] < 4 {
		if val, found := cache[start_idx+3]; found {
			c += val
		} else {
			result := GetAllCombos(data, start_idx+3, cache)
			cache[start_idx+3] = result
			c += result
		}
	}

	return c
}

func day10() {
	inp, _ := ioutil.ReadFile("./inputs/day10.input")

	data := GetIntInput(inp)
	// append 0 for part 2
	data = append(data, 0)
	sort.Ints(data)

	current_rating := 0
	diff1 := 0
	diff3 := 1
	for _, n := range data {
		if n-current_rating == 1 {
			diff1++
		} else if n-current_rating == 3 {
			diff3++
		}

		current_rating = n
	}

	fmt.Println("difs:", diff1, diff3)
	fmt.Println("part1:", diff1*diff3)

	cache := make(map[int]int)
	// too low: 7 895 290 740 736
	// too low: 7 895 290 740 736, tried it again on accident :(

	// failed part2 badly because sample input was not including zero;
	// zero->1, zero->2, zero->3 is multiple paths.
	// i was not factoring those in
	all_combos := GetAllCombos(data, 0, cache)
	fmt.Println("all combos:", all_combos)
}
