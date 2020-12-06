package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func day06() {
	inp, _ := ioutil.ReadFile("./day06.input")

	data := GetDoubleStringInput(inp)

	total_answers_any_yes := 0
	total_answers_all_yes := 0
	for _, answers := range data {
		answer_set := make(map[rune]int)
		for _, q := range answers {
			if q != 10 && q != 13 {
				answer_set[q] += 1
			}
		}

		total_answers_any_yes += len(answer_set)
		total_people := len(strings.Split(answers, "\r\n"))

		for _, val := range answer_set {
			if val == total_people {
				total_answers_all_yes++
			}
		}
	}

	fmt.Println("total answers anyone yes: ", total_answers_any_yes)
	fmt.Println("total answers all yes: ", total_answers_all_yes)
}
