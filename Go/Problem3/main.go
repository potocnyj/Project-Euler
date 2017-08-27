package main

import (
	"fmt"
	"math"
)

/*
	Problem 3: Largest prime factor.
	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?
*/
func main() {
	const limit = 600851475143
	// We want the largest prime factor of our limit;
	// take advantage of the fact that the largest possible
	// factor is sqrt(limit), and use that as an upper bound
	// for our prime sieve.
	sqLimit := int(math.Sqrt(float64(limit))) + 1
	seq := sieve(sqLimit)

	max := 0
	for n := range seq {
		if limit%n == 0 {
			max = n
		}
	}

	fmt.Printf("the largest prime factor of input %d is %d\n", limit, max)
}

// Implement the sieve of Eratosthenes via daisy-chained goroutines.
func sieve(limit int) <-chan int {
	ch := generate(limit)
	res := make(chan int)
	go func() {
		for {
			prime, ok := <-ch
			if !ok {
				break
			}
			res <- prime

			ch1 := make(chan int)
			go filter(ch, ch1, prime)
			ch = ch1
		}

		close(res)
	}()

	return res
}

// generate returns a channel to receive the sequence 2,3,4... through.
func generate(limit int) <-chan int {
	res := make(chan int)

	go func() {
		for i := 2; i < limit; i++ {
			res <- i
		}
		close(res)
	}()

	return res
}

// filter passes values that are not a multiple of prime
// through the input channel pipeline.
func filter(in <-chan int, out chan<- int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}
