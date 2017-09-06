package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	var input int64 = 9009
	if pal := isPalindrome(input); !pal {
		t.Errorf("expected %d to be a palindrome, got %v instead", input, pal)
	}

	input = 123
	if pal := isPalindrome(input); pal {
		t.Errorf("expected %d to NOT be a palindrome, got %v instead", input, pal)
	}

	input = 12321
	if pal := isPalindrome(input); !pal {
		t.Errorf("expected %d to be a palindrome, got %v instead", input, pal)
	}

	input = 123321
	if pal := isPalindrome(input); !pal {
		t.Errorf("expected %d to be a palindrome, got %v instead", input, pal)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if !isPalindrome(9009) {
			b.Fail()
		}
	}
}

func TestGetLargestPalindromeProduct(t *testing.T) {
	terms := []int64{}
	for i := int64(10); i < 100; i++ {
		terms = append(terms, i)
	}

	pal := getLargestPalindromeProduct(terms)
	if expected := int64(9009); pal != expected {
		t.Errorf("expected output %d, got %d instead", expected, pal)
	}
}

func BenchmarkGetLargestPalindromeProduct(b *testing.B) {
	terms := []int64{}
	for i := int64(10); i < 100; i++ {
		terms = append(terms, i)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if getLargestPalindromeProduct(terms) == 0 {
			b.Fail()
		}
	}
}
