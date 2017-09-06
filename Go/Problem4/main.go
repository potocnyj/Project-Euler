package main

import (
	"fmt"
	"strconv"
)

/*
	Problem 4: Largest palindrome product.
	A palindromic number reads the same both ways.
	The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

	Find the largest palindrome made from the product of two 3-digit numbers.
*/
func main() {
	terms := make([]int64, 900)
	for i := int64(0); i < 900; i++ {
		terms[i] = i + 100
	}

	pal := getLargestPalindromeProduct(terms)
	fmt.Println("The largest palindrome made from the product of two 3-digit numbers is", pal)
}

// getLargestPalindromeProduct iterates through the input list of terms
// and multiplies them against each other, looking for products that
// are a palindrome (e.g. 9009).
// It returns the largest palindrome found.
func getLargestPalindromeProduct(terms []int64) int64 {
	var max int64
	for ix, x := range terms {
		if ix == len(terms) {
			break
		}

		for _, y := range terms[ix+1:] {
			prod := x * y
			if isPalindrome(prod) && prod > max {
				max = prod
			}
		}
	}

	return max
}

// isPalindrome returns true if the input number is a palindrome, e.g. 9009.
func isPalindrome(x int64) bool {
	xstr := strconv.FormatInt(x, 10)

	pivot := len(xstr) / 2
	for i := 0; i < pivot; i++ {
		if xstr[i] != xstr[len(xstr)-1-i] {
			return false
		}
	}

	return true
}
