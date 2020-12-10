package main

import (
	"fmt"
	"io/ioutil"
)

// In your batch file, how many passports are valid
func day04() {
	inp, _ := ioutil.ReadFile("./inputs/day04.input")

	data := GetDoubleStringInput(inp)

	count := 0
	for _, raw_passport := range data {
		_, err := New(raw_passport)

		if err == nil {
			count++
		}
	}

	fmt.Println("valid passports: ", count)
}
