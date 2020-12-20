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

func CheckInputWithRules(inp string, rules map[int]Rule, rule_num int) (bool, string) {
	rule := rules[rule_num]

	if rule.Letter != "" {
		if debug {
			fmt.Printf("got a letter [%s] == %s\n", rule.Letter, inp)
		}
		if len(inp) == 0 {
			return false, "-1"
		}
		return string(inp[0]) == rule.Letter, inp[1:]
	}

	mixes := [][]int{rule.Mix1, rule.Mix2}
	good := ""
	for i, mix := range mixes {
		if debug {
			fmt.Printf("\n\ntrying mix[%d]: %v\n", i, mix)
		}
		if len(mix) > 0 {
			checked_input := inp
			matched := false

			for _, k := range mix {
				if debug {
					fmt.Printf("checking if \"%s\" input matches rule [%d]: %v\n", checked_input, k, rules[k])
				}
				matched, checked_input = CheckInputWithRules(checked_input, rules, k)
				if debug {
					fmt.Printf("matched: %v, with rest: %s\n", matched, checked_input)
				}
				if !matched {
					break
				}
			}

			if debug {
				fmt.Println("==got to end of loop with rest:", checked_input)
			}
			if matched && checked_input == "" {
				return true, ""
			}
			if matched {
				if debug {
					fmt.Println("good ->", checked_input)
				}
				good = checked_input
				return true, checked_input
			}
		}
	}

	if debug {
		fmt.Println("===== at end of whole thing with good:", good)
	}

	return false, good
}

func day19() {
	inp, _ := ioutil.ReadFile("./inputs/day19.sample")

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
	for _, k := range inputs {
		matched, rest := CheckInputWithRules(k, rules, 0)
		fmt.Println("final [", k, "] matches:", matched, "rest:", rest)
		if matched && rest == "" {
			c++
		}
	}
	fmt.Println("total matches:", c)
}
