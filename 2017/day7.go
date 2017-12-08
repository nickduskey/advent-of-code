package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getWeightMap() map[string]int {
	i, _ := ioutil.ReadFile("day7.txt")

	input := string(i)
	rows := strings.Split(input, "\n")

	m := make(map[string]int)

	for _, row := range rows {
		a := strings.Split(row, "(")
		b := strings.Split(a[1], ")")
		int, _ := strconv.Atoi(b[0])
		m[a[0]] = int
	}

	return m
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func day7p1() {
	i, _ := ioutil.ReadFile("day7.txt")

	input := string(i)
	rows := strings.Split(input, "\n")

	var names, heldup []string
	for _, row := range rows {
		a := strings.Split(row, "(")
		name := strings.TrimSpace(a[0])
		names = append(names, name)
		b := strings.Split(a[1], "-> ")
		if len(b) == 1 {
			continue
		}
		c := strings.Split(b[1], ",")
		for _, s := range c {
			n := strings.TrimSpace(s)
			heldup = append(heldup, n)
		}
	}
	// Find the name that isn't in the heldUp
	for _, name := range names {
		b := stringInSlice(name, heldup)
		if b == false {
			fmt.Println("pt1 answer:", name)
		}
	}
}

func main() {
	day7p1()
}
