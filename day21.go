package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func ArrayContains(arr []string, value string) bool {
	for _, k := range arr {
		if k == value {
			return true
		}
	}
	return false
}

func FindAllergen(foods []Food, allergen string) string {
	foods_that_contain := []Food{}
	for _, food := range foods {
		if ArrayContains(food.Allergens, allergen) {
			foods_that_contain = append(foods_that_contain, food)
		}
	}

	counts := make(map[string]int)
	for _, food := range foods_that_contain {
		for _, ing := range food.Ingredients {
			counts[ing]++
		}
	}

	counts_that_match := 0
	name_that_match := ""
	for name, count := range counts {
		if string(name[0]) != "!" && count == len(foods_that_contain) {
			counts_that_match++
			name_that_match = name
		}
	}

	if counts_that_match > 1 {
		return ""
	}
	return name_that_match
}

func AdjustNames(foods []Food, allergen string, real_name string) []Food {
	new_foods := []Food{}

	for _, food := range foods {
		if ArrayContains(food.Ingredients, real_name) {
			for idx, ing := range food.Ingredients {
				if ing == real_name {
					// weird hack the FindAllergen function knows that this is found already
					food.Ingredients[idx] = "!" + allergen
				}
			}
		}
		new_foods = append(new_foods, food)
	}

	return new_foods
}

func day21() {
	inp, _ := ioutil.ReadFile("./inputs/day21.input")

	data := GetStringInput(inp)

	allergens := make(map[string]string)
	ingredients := make(map[string]string)
	foods := []Food{}

	for _, line := range data {
		contains_split := strings.Split(line, " (contains ")
		ings := strings.Split(contains_split[0], " ")
		als := strings.Split(strings.Trim(contains_split[1], ")"), ", ")

		for _, a := range als {
			allergens[a] = "unknown"
		}
		for _, i := range ings {
			ingredients[i] = "unknown"
		}

		foods = append(foods, Food{ings, als})
	}

	fmt.Println("foods:", len(foods), "i:", len(ingredients), "a:", len(allergens))
	fmt.Println(allergens)

	still_going := true
	for still_going {
		still_going = false
		for a, _ := range allergens {
			real_name := FindAllergen(foods, a)

			if real_name != "" {
				foods = AdjustNames(foods, a, real_name)
				fmt.Println("setting", a, "to", real_name)
				allergens[a] = real_name
				still_going = true
			}
		}
	}

	count_non_allergens := 0
	for _, food := range foods {
		for _, ings := range food.Ingredients {
			if string(ings[0]) != "!" {
				count_non_allergens++
			}
		}
	}

	fmt.Println(count_non_allergens)

	// gather all the allergens to sort them later
	algs := []string{}
	for k := range allergens {
		algs = append(algs, k)
	}

	sort.Strings(algs)
	for _, ag := range algs {
		fmt.Printf("%v,", allergens[ag])
	}
}
