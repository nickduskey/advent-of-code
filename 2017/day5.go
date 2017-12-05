package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() []int {
	bytes, err := ioutil.ReadFile("day5input.txt")
	if err != nil {
		panic(err)
	}

	str := string(bytes)
	strs := strings.Split(str, "\n")
	var prepared []int
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		prepared = append(prepared, i)
	}
	return prepared
}

func day5part1(input []int) int {
	i := 0
	steps := 0
	// while the index is within the instruction set
	for {
		steps++
		v := input[i]
		input[i]++
		i += v
		if i >= len(input) || i < 0 {
			return steps
		}
	}
}

func day5part2(input []int) int {
	i := 0
	steps := 0
	// while the index is within the instruction set
	for {
		steps++
		v := input[i]
		if v >= 3 {
			input[i]--
		} else {
			input[i]++
		}
		i += v
		if i >= len(input) || i < 0 {
			return steps
		}
	}
}

func main() {
	prepared := getInput()
	steps := day5part1(prepared)
	p2 := day5part2(prepared)
	fmt.Println("p1:", steps)
	fmt.Println("p2:", p2)
}
