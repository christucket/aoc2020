package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func GetPasswordRules(password string) (int, int, string) {
	rules := strings.Split(password, ":")
	times_letter := strings.Split(rules[0], " ")
	letter := times_letter[1]
	at_least, _ := strconv.Atoi(strings.Split(times_letter[0], "-")[0])
	at_most, _ := strconv.Atoi(strings.Split(times_letter[0], "-")[1])

	return at_least, at_most, letter
}

func day02() {
	inp, _ := ioutil.ReadFile("./inputs/day02.input")

	data := GetStringInput(inp)

	sled_valid := 0
	toboggan_valid := 0

	for _, line := range data {
		password := strings.Split(line, ": ")[1]
		at_least, at_most, letter := GetPasswordRules(line)

		count := strings.Count(password, letter)
		if count >= at_least && count <= at_most {
			sled_valid++
		}

		first_match := len(password) >= at_least && string(password[at_least-1]) == letter
		second_match := len(password) >= at_most && string(password[at_most-1]) == letter
		if (first_match && !second_match) || (!first_match && second_match) {
			toboggan_valid++
		}
	}

	fmt.Println("sled_valid passwords: ", sled_valid)
	fmt.Println("toboggan_valid passwords: ", toboggan_valid)
}
