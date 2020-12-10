package main

import (
	"fmt"
	"io/ioutil"
)

func GetInvalidNumSumRange(data []int, looking_for int) (int, int) {
	length := 2
	for length < len(data) {
		start_at := 0
		// idx := start_at

		for start_at < len(data) {
			c := 0

			sum := 0
			// start sum loop from start...start+lengthh
			for c < length {
				sum += data[start_at+c]
				if sum > looking_for {
					break
				}
				c++
			}

			if sum == looking_for {
				fmt.Println("found??")
				fmt.Println("adding length", length, "starting at", start_at, "sum:", sum)
				return start_at, start_at + length - 1
			}

			start_at++
		}
		length++

	}
	return 0, 0
}

func day09() {
	inp, _ := ioutil.ReadFile("./day09.input")

	data := GetIntInput(inp)

	preample_size := 25

	i := preample_size
	invalid_num := 0

	for i < len(data) {
		num := data[i]
		available_range := data[i-preample_size : i]

		// check if sum exists
		found := false
		for _, ar_num1 := range available_range {
			for _, ar_num2 := range available_range {
				if ar_num1+ar_num2 == num && ar_num1 != ar_num2 {
					found = true
				}
			}
		}

		// fmt.Println(available_range)
		if !found {
			invalid_num = num
			break
		}

		i++
	}
	fmt.Println("invalid number:", invalid_num)

	low_range, high_range := GetInvalidNumSumRange(data, invalid_num)
	fmt.Println("sum range:", low_range, high_range)

	min := -1
	max := -1
	for low_range < high_range {
		num := data[low_range]
		if min == -1 || num < min {
			min = num
		}
		if max == -1 || num > max {
			max = num
		}
		low_range++
	}
	fmt.Println("min, max:", min, max, "added: ", min+max)

}
