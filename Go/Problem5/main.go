package main

import (
	"fmt"
	"math"
)

/*
	Problem 5: Smallest multiple.
	2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func main() {
	terms := make([]int64, 0, 20)
	for i := int64(1); i < 21; i++ {
		terms = append(terms, i)
	}

	prod := getSmallestCommonMultiple(terms)
	fmt.Println("The smallest multiple of the numbers from 1 to 20 is", prod)
}

// getSmallestCommonMultiple takes the input list of factors
// and identifies lowest number that is a multiple of all of the factors listed.
func getSmallestCommonMultiple(factors []int64) int64 {
	// Take the input list of factors and remove any
	// compound factors that can be composed by the others in the list.
	multiples := reduceFactors(factors)

	// Now iterate the list of multiples remaining
	// and multiply them together to get a resultant product.
	var product int64 = 1
	for _, x := range multiples {
		product *= x
	}

	return product
}

// reduceFactors takes the input list of terms and reduces
// it down to a list of prime factors to multiply together
// to get the Least Common Multiple.
func reduceFactors(terms []int64) []int64 {
	commonFactors := make(map[int64]int)
	for _, x := range terms {
		factors := getPrimeFactorization(x)
		for x, n := range factors {
			if commonFactors[x] < n {
				commonFactors[x] = n
			}
		}
	}

	distTerms := []int64{}
	for x, n := range commonFactors {
		for i := 0; i < n; i++ {
			distTerms = append(distTerms, x)
		}
	}

	return distTerms
}

// getPrimeFactorization returns the prime-factorization of x,
// expressed as a mapping of each prime factor
// to the number of times it divides x.
func getPrimeFactorization(x int64) map[int64]int {
	limit := int64(math.Sqrt(float64(x))) + 1

	results := make(map[int64]int)
	for x > 1 {
		hasMore := false
		for i := int64(2); i < limit; i++ {
			if x%i == 0 {
				// i is a factor of x.
				results[i]++
				x /= i
				hasMore = true
				break
			}
		}

		if !hasMore {
			// Looping again won't get us anywhere,
			// add the remaining value of x to results and return.
			results[x]++
			break
		}
	}

	return results
}
