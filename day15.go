package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func FindLastIndex(numbers []int, number int) int {
	// actually want to skip the first occurence
	for i := len(numbers) - 2; i >= 0; i-- {
		if numbers[i] == number {
			return i
		}
	}
	// should never happen in our case
	return -1
}

func day15() {
	inp, _ := ioutil.ReadFile("./inputs/day15.input")

	data := strings.Split(GetStringInput(inp)[0], ",")

	memory := make(map[int]int)

	idx := 0
	for i, v := range data {
		v, _ := strconv.Atoi(v)
		memory[v] = i
		idx = i
	}

	idx++

	// done at 5am, pretty sure i was making this over complicated with the next memory thing.
	// probably just need to reverse the if condition or something but i was adding the value to
	// memory instantly and it kept being found... obviously, since i was adding it immediately.
	next_memory_val := -1
	next_memory_idx := -1

	for idx < 30000000 {
		last_number := next_memory_idx

		v, f := memory[last_number]

		if next_memory_val != -1 {
			memory[next_memory_idx] = next_memory_val
		}

		if !f {
			next_memory_idx = 0
			next_memory_val = idx
		} else {
			num_diff := idx - 1 - v
			next_memory_idx = num_diff
			next_memory_val = idx
		}

		idx++
		if idx == 2020 {
			fmt.Println("last number spoken at", idx, " was:", next_memory_idx)
		}
	}

	fmt.Println("last number spoken at", idx, " was:", next_memory_idx)
}
