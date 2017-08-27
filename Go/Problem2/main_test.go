package main

import (
	"testing"
)

func TestGetFibTerms(t *testing.T) {
	res := getFibTerms(50)

	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	i := 0
	for n := range res {
		if i > len(expected) {
			t.Fatalf("received too many terms(i=%d)", i)
		}

		if n != expected[i] {
			t.Errorf("expected fib term %d to be %d, got %d instead.", i, expected[i], n)
		}

		i++
	}
}

func TestIsEven(t *testing.T) {
	cases := map[int]bool{
		1: false,
		2: true,
		5: false,
	}

	for in, expected := range cases {
		if isEven(in) != expected {
			t.Errorf("expected %d to be %t, got %t", in, expected, isEven(in))
		}
	}
}

var even bool

func BenchmarkIsEven(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		even = isEven(i)
	}
}

func BenchmarkGetFibTermsEven(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := getFibTermsFilter(4000000, isEven)
		for _ = range res {
		} // drain channel
	}
}
