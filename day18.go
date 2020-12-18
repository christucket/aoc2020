package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func DoOperator(operator string, ans int, val int) int {
	if operator == "+" {
		return ans + val
	} else if operator == "*" {
		return ans * val
	} else {
		DebugPrint("got [", ans, "] (", operator, ") [", val, "]")
	}

	return -5
}

func SolveEq(eq string, starting_operator string) (int, string) {
	first := true
	ans := 0
	operator := ""

	for i := 0; i < len(eq); i++ {
		c := string(eq[i])

		if first && c != "(" {
			ans = ToInt(c)
		} else if c == "(" {
			result, rest := SolveEq(eq[i+1:], "(")

			// reset loop
			i = -1
			eq = rest

			if operator == "" {
				ans = result
			} else {
				ans = DoOperator(operator, ans, result)
			}
		} else if c == ")" {
			if starting_operator == "(" {
				return ans, eq[i+1:]
			} else if starting_operator == "*" {
				return ans, eq[i:]
			}
		} else if c == "*" { /* treat it like a special paren */
			result, rest := SolveEq(eq[i+1:], "*")

			// reset loop
			i = -1
			eq = rest

			ans = DoOperator("*", ans, result)
		} else if c == "+" {
			operator = c
		} else /* character is digit? */ {
			ans = DoOperator(operator, ans, ToInt(c))
		}

		first = false
	}

	return ans, ""
}

// try to put parens around everything but multi
func VerboseEq(eq string) string {
	split := strings.SplitN(eq, "*", 2)

	if len(split) > 1 {
		a, b := split[0], split[1]
		if len(a) >= 0 {
			a = "(" + a + ")"
		}
		if len(b) >= 0 {
			b = "(" + b + ")"
		}
		verbose := "" + a + "*" + VerboseEq(b) + ""
		return verbose
	} else {
		return eq
	}
}

func day18() {
	inp, _ := ioutil.ReadFile("./inputs/day18.input")

	data := GetStringInput(inp)

	sum := 0
	for _, l := range data {
		l = strings.ReplaceAll(l, " ", "")
		ans, rest := SolveEq(l, "")

		if rest != "" {
			fmt.Println("broken eq", l)
		} else {
			sum += ans
		}
	}
	fmt.Println(sum)
}
