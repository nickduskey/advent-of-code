package main

import "fmt"

func day6p1() int {
	banks := [16]int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	// fmt.Println("starting banks:", banks)
	var cycles, largest, largestIndex int
	var states [][16]int
	states = append(states, banks)
	// fmt.Println("starting states:", states)
	for {
		largest = 0
		largestIndex = 0
		// Get the largest bank
		for i, bank := range banks {
			if bank > largest {
				largest = bank
				largestIndex = i
			}
		}
		// Distribute the largest bank evenly
		q := largest / len(banks)
		r := largest % len(banks)

		// Handle q that is less than len(banks)
		if q == 0 {
			banks[largestIndex] = 0
			for i := 0; r > 0; i++ {
				if i == largestIndex {
					continue
				}
				banks[i]++
				r--
			}
		}
		// if even
		if r == 0 {
			banks[largestIndex] = 0
			// add q to all positions of banks
			for i := 0; i < len(banks); i++ {
				banks[i] += q
			}
		}
		// if remainder
		if r != 0 {
			fmt.Println("q:", q)
			fmt.Println("r:", r)
			// first set largestIndex to 0
			banks[largestIndex] = 0

			// give uneven to the largestIndex
			for i := 0; i < len(banks); i++ {
				if i == largestIndex {
					banks[i] += r
				} else {
					banks[i] += q
				}
			}
			// fmt.Println("banks after modification:", banks)
		}

		// Check the modified banks against the previously seen states
		for i, state := range states {
			if banks == state {
				return cycles
			}
		}

		// If unseen save the current banks as a new state in states
		states = append(states, banks)

		// Increment the cycle count
		cycles++
	}
}

func main() {
	cycles := day6p1()
	fmt.Println("cycles:", cycles)
}
