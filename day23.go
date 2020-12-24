package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Node struct {
	val      int
	NextNode *Node
}

var size int
var cache map[int]*Node

func FindDestinationCup(current_cup *Node, taken_out_cups *Node) *Node {
	looking_for := current_cup.val - 1
	if looking_for == 0 {
		looking_for = size
	}
	restart := true

	for restart {
		restart = false
		looking_at := taken_out_cups
		for range []int{1, 2, 3} {
			if looking_at.val == looking_for {
				looking_for--
				if looking_for == 0 {
					looking_for = size
				}
				restart = true
				break
			}
			looking_at = looking_at.NextNode
		}
	}

	return cache[looking_for]
}

func GetRingString(current_cup *Node, sep string) string {
	f := ""
	first := current_cup

	for {
		f += fmt.Sprintf("%v"+sep, current_cup.val)
		current_cup = current_cup.NextNode

		if current_cup == first {
			break
		}
	}

	return f[:len(f)-len(sep)]
}

func PlayCupsGame(ring *Node, number_of_moves int) *Node {
	current_cup := ring

	move := 1
	for move <= number_of_moves {
		next_three_cups := current_cup.NextNode

		// set the nextnode to 3 values after the current cup
		// takes 3 cups out of the loop
		current_cup.NextNode = current_cup.NextNode.NextNode.NextNode.NextNode
		// find the destination cup, has to use three cups to avoid using those values
		destination_cup := FindDestinationCup(current_cup, next_three_cups)

		// need to store the value after the destination cup so
		// we can set it as the NextNode to the last of the 3 cups
		after_destination := destination_cup.NextNode
		// set the next code after destination cup to the three cups in order
		destination_cup.NextNode = next_three_cups
		next_three_cups.NextNode.NextNode.NextNode = after_destination

		// next cup should be (by rules) the next cup in line
		current_cup = current_cup.NextNode

		move++
	}

	return ring
}

func day23() {
	inp, _ := ioutil.ReadFile("./inputs/day23.input")

	data := GetStringInput(inp)

	part2 := true
	cache = make(map[int]*Node)

	var previous_node *Node
	var first_node *Node

	numbers := []int{}
	for _, k := range data[0] {
		next_int, _ := strconv.Atoi(string(k))
		numbers = append(numbers, next_int)
	}

	if part2 {
		for i := 10; i <= 1_000_000; i++ {
			numbers = append(numbers, i)
		}
	}
	for _, k := range numbers {
		size++

		node := Node{k, first_node}
		cache[k] = &node

		if first_node == nil {
			first_node = &node
			first_node.NextNode = &node
		} else {
			previous_node.NextNode = &node
		}
		previous_node = &node
	}

	total_moves := 100
	if part2 {
		total_moves = 10_000_000
	}
	final_ring := PlayCupsGame(first_node, total_moves)

	// find the cup with the label 1
	one_node := final_ring
	for {
		if one_node.val == 1 {
			break
		}
		one_node = one_node.NextNode
	}

	if part2 {
		fmt.Println("final:", one_node, one_node.NextNode, one_node.NextNode.NextNode)
		fmt.Println("prod:", one_node.NextNode.val*one_node.NextNode.NextNode.val)
	} else {
		final_string := GetRingString(one_node.NextNode, "")
		final_string = final_string[:len(final_string)-1]
		fmt.Println(final_string)
	}
}
