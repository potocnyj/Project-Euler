package main

import (
	"testing"
)

func TestSieve(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	out := sieve(30)
	i := 0
	for num := range out {
		if i == len(expected) {
			t.Fatalf("received more primes than expected: %d", i)
		}

		if expected[i] != num {
			t.Errorf("expected output %d to be %d, got %d instead.", i, expected[i], num)
		}

		i++
	}
}
