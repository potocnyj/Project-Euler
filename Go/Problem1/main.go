package main

import (
	"fmt"
)

const problemLimit = 1000

/*
	Problem 1:
	If we list all the natural numbers below 10
	that are multiples of 3 or 5, we get 3, 5, 6 and 9.
	The sum of these multiples is 23.

	Find the sum of all the multiples of 3 or 5 below 1000.
*/
func main() {
	fmt.Printf("Sum of multiples of 3 or 5 less than %d: %d\n",
		problemLimit, sumInts(getMultiples(problemLimit, []int{3, 5}...)))
}

// getMultiples returns a list of integers that are a multiple
// of one the list of provided divisors, up to the specified limit.
func getMultiples(limit int, divisors ...int) []int {
	if len(divisors) == 0 {
		// Just return if no divisors are provided.
		return nil
	}

	// Keep things simple, iterate up to limit and
	// check divisors to see if we have a multiple as we go.
	//
	// TODO - make this faster? It runs quickly enough but I can probably do better.
	var multiples []int
	for i := 1; i < limit; i++ {
		isMultiple := false
		for _, div := range divisors {
			if rem := i % div; rem == 0 {
				isMultiple = true
				break
			}
		}

		if isMultiple {
			multiples = append(multiples, i)
		}
	}

	return multiples
}

// sumInts adds up a list of integers and returns the total.
func sumInts(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
