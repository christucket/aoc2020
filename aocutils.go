package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetIntInput(raw_data []byte) []int {
	split_data := strings.Split(string(raw_data), "\r\n")
	// convert string[] to int[]
	var data = []int{}
	for _, i := range split_data {
		j, _ := strconv.Atoi(i)
		data = append(data, j)
	}

	return data
}

func GetStringInput(raw_data []byte) []string {
	split_data := strings.Split(string(raw_data), "\r\n")

	return split_data
}

func GetDoubleStringInput(raw_data []byte) []string {
	split_data := strings.Split(string(raw_data), "\r\n\r\n")

	return split_data
}

func Get2DStringInput(raw_data []byte) [][]string {
	split_data := strings.Split(string(raw_data), "\r\n")

	var ret [][]string

	for _, s := range split_data {
		ret = append(ret, strings.Split(s, ""))
	}

	return ret
}

func DebugPrint(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}
