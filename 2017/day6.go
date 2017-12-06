package main

import "fmt"

func day6p1() (int, [16]int) {
	banks := [16]int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	var cycles, largest, largestIndex int
	var states [][16]int
	states = append(states, banks)
	for {
		// Increment the cycle count
		cycles++

		largest = 0
		largestIndex = 0
		// Get the largest bank
		for i, bank := range banks {
			if bank > largest {
				largest = bank
				largestIndex = i
			}
		}

		val := largest

		banks[largestIndex] = 0
		i := largestIndex + 1
		if i > (len(banks) - 1) {
			i = 0
		}

		for val > 0 {
			banks[i]++
			i++
			if i > (len(banks) - 1) {
				i = 0
			}
			val--
		}

		// Check the modified banks against the previously seen states
		for _, state := range states {
			if banks == state {
				return cycles, banks
			}
		}

		// If unseen save the current banks as a new state in states
		states = append(states, banks)
	}
}

func day6p2(banks [16]int) int {
	var cycles, largest, largestIndex int
	var states [][16]int
	states = append(states, banks)
	for {
		// Increment the cycle count
		cycles++

		largest = 0
		largestIndex = 0
		// Get the largest bank
		for i, bank := range banks {
			if bank > largest {
				largest = bank
				largestIndex = i
			}
		}

		val := largest

		banks[largestIndex] = 0
		i := largestIndex + 1
		if i > (len(banks) - 1) {
			i = 0
		}

		for val > 0 {
			banks[i]++
			i++
			if i > (len(banks) - 1) {
				i = 0
			}
			val--
		}

		// Check the modified banks against the previously seen states
		for _, state := range states {
			if banks == state {
				return cycles
			}
		}

		// If unseen save the current banks as a new state in states
		states = append(states, banks)
	}
}

func main() {
	cycles, banks := day6p1()
	addlcycles := day6p2(banks)
	fmt.Println("cycles:", cycles)
	fmt.Println("addlcycles:", addlcycles)
}
