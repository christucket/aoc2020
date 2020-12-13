package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func day13() {
	inp, _ := ioutil.ReadFile("./inputs/day13.input")

	data := GetStringInput(inp)

	departure_time, _ := strconv.Atoi(data[0])
	strnums := strings.Split(data[1], ",")

	var nums []int
	for _, n := range strnums {
		intnum, _ := strconv.Atoi(n)
		nums = append(nums, intnum)
	}

	target_bus, min_diff := 0, 1000
	for _, n := range nums {
		// x * num >= departure_time
		if n == 0 {
			continue
		}
		x := float64(departure_time) / float64(n)
		if int(math.Round(x+0.5))*n-departure_time < min_diff {
			target_bus = n
			min_diff = int(math.Round(x+0.5))*n - departure_time
		}
		fmt.Println("bus #", n, ":   ", x, "loops, ", int(math.Round(x+0.5))*n, (int(x)+1)*n)
	}

	fmt.Println("yep", departure_time, "bus: ", target_bus, ": time: ", min_diff, target_bus*min_diff)

	// try brute force to see if go is fast enough
	// pretty fast but probably not fast enough for main input.
	// brute forced the hell out of this anyways with a slight optimization after looking at the bus data

	// t, diff, jumps := 498597000000000, 29, 15109
	// t, diff, jumps := 536309064000000, 29, 15109
	t, diff, jumps := 722149764000000, 29, 15109
	//t, diff, jumps := 133, 7, 133
	loops := 0
	for {
		found := true

		for i, n := range nums {
			if n == 0 {
				continue
			}
			if (t+i-diff)%n > 0 {
				found = false
				break
			}
		}
		if found {
			fmt.Println("found good time at:", t-diff)
			break
		}

		t += jumps
		loops += 1
		if loops%100000000 == 0 {
			fmt.Println("checking t...", t, loops)
		}
	}

	// looking at the numbers,
	// sample has bus  7 at t0
	// and        bus 19 at t7
	// we have bus 41  at t19
	//     and bus 601 at t60, which means bus 41 and bus 601 will collide
	// so 41*601*n == time+61
	// and
	// we have bus 29  at t0
	//     and bus 521 at t29, which means we have two sets here

}
