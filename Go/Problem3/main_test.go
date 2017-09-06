package main

import (
	"reflect"
	"testing"
)

func TestGetPrimeFactors(t *testing.T) {
	// The prime factors of 13195 are 5, 7, 13 and 29.
	expected := []int64{5, 7, 13, 29}
	factors := getPrimeFactors(13195)

	if !reflect.DeepEqual(factors, expected) {
		t.Errorf("got factors %v\nexpected %v", factors, expected)
	}
}

func BenchmarkGetPrimeFactors(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		getPrimeFactors(13195)
	}
}
