package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// i don't know any of the syntax conventions
func find_dups(data []int) (int, int) {
	// find the two entries that sum to 2020
	for _, num1 := range data {
		if num1 > 2020 {
			continue
		}
		for _, num2 := range data {
			if num1+num2 == 2020 {
				return num1, num2
			}
		}
	}

	return -1, -1
}

func find_trios(data []int) (int, int, int) {
	// find the two entries that sum to 2020
	for _, num1 := range data {
		if num1 > 2020 {
			continue
		}
		for _, num2 := range data {
			if num1+num2 > 2020 {
				continue
			}
			for _, num3 := range data {
				if num1+num2+num3 == 2020 {
					return num1, num2, num3
				}
			}
		}
	}

	return -1, -1, -1
}

func main() {
	inp, _ := ioutil.ReadFile("./day01.input")

	raw_data := strings.Split(string(inp), "\r\n")
	// convert string[] to int[]
	var data = []int{}
	for _, i := range raw_data {
		j, _ := strconv.Atoi(i)
		data = append(data, j)
	}

	i, j := find_dups(data)
	fmt.Printf("solution %d+%d = %d\t\tx %d\n", i, j, i+j, i*j)

	i, j, k := find_trios(data)
	fmt.Printf("solution %d+%d+%d = %d\t\tx %d", i, j, k, i+j+k, i*j*k)
}
