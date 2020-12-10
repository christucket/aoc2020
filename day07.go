package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func GetBags(bag_rule string) (string, map[string]string) {
	children := make(map[string]string)

	regexp_pattern := "(.*?) bags contain "
	pattern := regexp.MustCompile(regexp_pattern).FindStringSubmatch(bag_rule)
	bag_color := pattern[1]

	if strings.Index(bag_rule, "no other bags") > 0 {
		return bag_color, children
	}

	regexp_pattern = " (\\d) (.*?) bags?"
	matcher := regexp.MustCompile(regexp_pattern).FindAllStringSubmatch(bag_rule, -1)

	for _, match := range matcher {
		children[match[2]] = match[1]
	}

	return bag_color, children
}

func GetAllParentBags(bag_map map[string]map[string]string, token_bag string) map[string]bool {
	possible_bags := make(map[string]bool)
	// find all direct bags
	for bag_color, children := range bag_map {
		if _, found := children[token_bag]; found {
			possible_bags[bag_color] = false
		}
	}

	c := 0
	// find in-direct bags using the current list
	for c != len(possible_bags) {
		c = len(possible_bags)

		// loop over all next possible bags
		for bag_color, checked := range possible_bags {
			if !checked {
				// set checked to true
				possible_bags[bag_color] = true

				// add parents to possible_bags
				for p, children := range bag_map {
					if _, found := children[bag_color]; found {
						possible_bags[p] = false
					}
				}
			}
		}
	}

	return possible_bags
}

func GetTotalChildren(bag_map map[string]map[string]string, token_bag string) int {
	total := 0
	for bag, child := range bag_map[token_bag] {
		int_child, _ := strconv.Atoi(child)
		total_children := GetTotalChildren(bag_map, bag)
		total = total + int_child + int_child*total_children
	}

	return total
}

// some kind of linked list thingy?
func day07() {
	inp, _ := ioutil.ReadFile("./inputs/day07.input")

	data := GetStringInput(inp)

	bag_map := make(map[string]map[string]string)

	for _, bag_rule := range data {
		bc, c := GetBags(bag_rule)
		bag_map[bc] = c
	}

	possible_parents := GetAllParentBags(bag_map, "shiny gold")
	total_children := GetTotalChildren(bag_map, "shiny gold")
	fmt.Println("all possible combos: ", len(possible_parents))
	fmt.Println("all children: ", total_children)
}
