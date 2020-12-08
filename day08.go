package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	Type     string
	Argument int
	Ran      bool
}

func ConvertDataIntoInstuctions(data []string) []Instruction {
	var instructions []Instruction

	for _, line := range data {
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])

		i := Instruction{op, arg, false}
		instructions = append(instructions, i)
	}

	return instructions
}

func ResetProgram(instructions []Instruction) {
	// need to figure out why/how i can do _, instruction here instead, pointer stuff?
	for i, _ := range instructions {
		instruction := &instructions[i]
		instruction.Ran = false
	}
}

func RunProgram(instructions []Instruction, swap int) (int, int, string) {
	accumulator := 0
	instruction_pointer := 0

	for !instructions[instruction_pointer].Ran {
		instruction := &instructions[instruction_pointer]

		// does go have switches?
		if instruction.Type == "acc" {
			accumulator += instruction.Argument
			instruction_pointer++
		} else if (instruction.Type == "jmp" && swap != instruction_pointer) ||
			(instruction.Type == "nop" && swap == instruction_pointer) {
			instruction_pointer += instruction.Argument
		} else {
			instruction_pointer++
		}

		// we ran the instruction
		instruction.Ran = true

		if instruction_pointer == len(instructions) {
			return instruction_pointer, accumulator, "graceful"
		}
		if instruction_pointer >= len(instructions) {
			return instruction_pointer, accumulator, "bad jmp"
		}
	}

	return instruction_pointer, accumulator, "infinite loop"
}

func day08() {
	inp, _ := ioutil.ReadFile("./day08.input")

	data := GetStringInput(inp)

	instructions := ConvertDataIntoInstuctions(data)

	_, infinite_acc, _ := RunProgram(instructions, -1)
	fmt.Println("No swaps acccumuator", infinite_acc)

	// going to brute force part 2
	for i, _ := range instructions {
		ResetProgram(instructions)
		ip, acc, status := RunProgram(instructions, i)

		if status == "graceful" {
			fmt.Println("\nending with pointer: ", ip)
			fmt.Println("            acc: ", acc)
			fmt.Println("            status: ", status)
			fmt.Println("            swap: ", i)
		}
	}

}
