package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func loop(subject_number int, loops int) int {
	output_num := 1

	for loops > 0 {
		output_num = (output_num * subject_number) % 20201227
		loops--
	}

	return output_num
}

func day25() {
	inp, _ := ioutil.ReadFile("./inputs/day25.input")

	data := GetStringInput(inp)

	card_public, _ := strconv.Atoi(data[0])
	door_public, _ := strconv.Atoi(data[1])

	card_loop_size := 0
	door_loop_size := 0

	loop_times := 100_000_000
	card_num := 1
	for i := 1; i < loop_times; i++ {
		card_num = (card_num * 7) % 20201227

		// fmt.Println(i, card_public, output_num, card_public == output_num)
		if card_public == card_num {
			card_loop_size = i
			break
		}
	}

	door_num := 1
	for i := 1; i < loop_times; i++ {
		door_num = (door_num * 7) % 20201227

		// fmt.Println(i, door_public, output_num, door_public == output_num)
		if door_public == door_num {
			door_loop_size = i
			break
		}
	}

	fmt.Println("card:", loop(door_public, card_loop_size))
	fmt.Println("card:", loop(card_public, door_loop_size))

	fmt.Println(card_loop_size, door_loop_size)
}
