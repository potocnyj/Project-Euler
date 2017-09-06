package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
	Problem 3: Largest prime factor.
	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?
*/
func main() {
	const limit = 600851475143

	factors := getPrimeFactors(limit)

	// Our prime factors are in ascending order,
	// so the last element of the slice
	// is the largest prime factor of our input value.
	max := factors[len(factors)-1]

	fmt.Printf("the largest prime factor of input %d is %d\n", limit, max)
}

func getPrimeFactors(N int64) []int64 {
	if N < 1 {
		return nil
	}

	// We want the list of prime factors of our input;
	// take advantage of the fact that the largest possible
	// factor is sqrt(N), and use that as an upper bound
	// for checking prime factors.
	//
	// For simplicity, just add one to the sqrt to account
	// for the likelihood that sqrt(N) isn't a factor.
	limit := int64(math.Sqrt(float64(N))) + 1
	var res []int64
	x := new(big.Int)
	for i := int64(2); i < limit; i++ {
		if N%i == 0 {
			// i is a factor of our input. Is it prime?
			x = x.SetInt64(i)
			if x.ProbablyPrime(20) {
				// It's prime, add it to the results.
				res = append(res, i)
			}
		}
	}

	return res
}
