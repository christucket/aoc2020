package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// this challenge is taken too literal. there is a way to create two masks for every mask
// and just apply the masks every time to do a & and |

// experiment with go funcs (generators??)
// counting from 1..n**2 where n is the number of floating bits, we get every permutation of bits
func GetMemoryAddresses(mask map[int]int, mem_idx int, c chan<- int) {
	loop := 1
	for _, v := range mask {
		if v == 2 {
			loop *= 2
		}
	}

	for loop > 0 {
		loop--
		mem_idx_floating := mem_idx
		loop_idx := 0
		for idx := 0; idx < len(mask); idx++ {
			v := mask[idx]
			if v == 2 {
				if loop&(1<<loop_idx) == 0 {
					mem_idx_floating &^= (1 << idx)
				} else {
					mem_idx_floating |= (1 << idx)
				}
				loop_idx += 1
			}
		}
		c <- mem_idx_floating
	}

	// we're done
	close(c)
}

func day14() {
	inp, _ := ioutil.ReadFile("./inputs/day14.input")

	data := GetStringInput(inp)

	registers := make(map[int]int)
	registers_p2 := make(map[int]int)
	mask := make(map[int]int)

	for _, line := range data {
		line_data := strings.Split(line, " = ")
		if line[0:4] == "mask" {
			mask_data := line_data[1]
			// clear mask (part 1 mistake 1)
			// too high: 2715539231000
			mask = make(map[int]int)
			for i, c := range mask_data {
				if string(c) == "X" {
					mask[len(mask_data)-i-1] = 2 // floating bit
				} else {
					c, _ := strconv.Atoi(string(c))
					mask[len(mask_data)-i-1] = c
				}
			}
		} else {
			mem_idx, _ := strconv.Atoi(strings.Split(strings.Split(string(line_data[0]), "[")[1], "]")[0])
			mem_idx_p2 := mem_idx
			mem_val_p1, _ := strconv.Atoi(line_data[1])
			mem_val := mem_val_p1

			// part 1
			for bit_pos, bit_val := range mask {
				if bit_val == 1 {
					mem_idx_p2 |= (1 << bit_pos)
				}

				if bit_val == 1 {
					mem_val_p1 |= (1 << bit_pos)
				} else if bit_val == 0 {
					mem_val_p1 &^= (1 << bit_pos)
				}
			}

			registers[mem_idx] = mem_val_p1

			// part 2
			mem_idx_chan := make(chan int)
			go GetMemoryAddresses(mask, mem_idx_p2, mem_idx_chan)

			for mem_idx_i := range mem_idx_chan {
				registers_p2[mem_idx_i] = mem_val
			}
		}
	}

	sum_p1 := 0
	for _, v := range registers {
		sum_p1 += v
	}

	sum_p2 := 0
	for _, v := range registers_p2 {
		sum_p2 += v
	}

	fmt.Println("part 1 sum: ", sum_p1)
	fmt.Println("part 2 sum: ", sum_p2)
}
