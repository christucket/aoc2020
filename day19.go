package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule struct {
	Mix1   []int
	Mix2   []int
	Letter string
}

func CheckInputWithRules(current_key int, next_keys []int, search string, index int, rules map[int]Rule) bool {
	rule := rules[current_key]

	if rule.Letter != "" {
		if index >= len(search) {
			return false
		}

		if rule.Letter == search[index:index+1] {
			if len(next_keys) > 0 {
				f0 := CheckInputWithRules(next_keys[0], next_keys[1:], search, index+1, rules)
				if f0 {
					return true
				}
			} else {
				if index == len(search)-1 {
					return true
				}
			}

		}
		return false
	}

	f1 := CheckInputWithRules(rule.Mix1[0], append(rule.Mix1[1:], next_keys...), search, index, rules)
	if f1 {
		return true
	}

	if len(rule.Mix2) > 0 {
		f2 := CheckInputWithRules(rule.Mix2[0], append(rule.Mix2[1:], next_keys...), search, index, rules)
		if f2 {
			return true
		}
	}
	return false
}

func day19() {
	inp, _ := ioutil.ReadFile("./inputs/day19.input")

	data := GetStringInput(inp)

	parsing_rules := true
	rules := make(map[int]Rule)
	inputs := []string{}

	for i, k := range data {
		if k == "" {
			parsing_rules = false
			inputs = data[i+1:]
			break
		}
		if parsing_rules {
			num, _ := strconv.Atoi(strings.Split(k, ":")[0])
			if strings.Count(k, "\"") > 0 {
				rules[num] = Rule{nil, nil, string(k[len(k)-2])}
			} else if strings.Count(k, "|") == 0 {
				mix1 := strings.Split(k, ": ")[1]
				rules[num] = Rule{StringToIntArray(mix1, " "), nil, ""}
			} else {
				vals := strings.Split(k, " | ")
				mix1 := strings.Split(vals[0], ": ")[1]
				mix2 := vals[1]
				rules[num] = Rule{StringToIntArray(mix1, " "), StringToIntArray(mix2, " "), ""}
			}
		}
	}

	c := 0
	for _, ababab := range inputs {
		f := CheckInputWithRules(rules[0].Mix1[0], rules[0].Mix1[1:], ababab, 0, rules)
		if f {
			c++
		}
	}
	fmt.Println("total matches:", c)
}
