package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPrimeFactorization(t *testing.T) {
	cases := []struct {
		x        int64
		expected map[int64]int
	}{
		{
			x:        16,
			expected: map[int64]int{2: 4},
		},
		{
			x:        2,
			expected: map[int64]int{2: 1},
		},
		{
			x:        21,
			expected: map[int64]int{3: 1, 7: 1},
		},
		{
			x:        30,
			expected: map[int64]int{2: 1, 3: 1, 5: 1},
		},
		{
			x:        64,
			expected: map[int64]int{2: 6},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			factors := getPrimeFactorization(c.x)

			if !reflect.DeepEqual(c.expected, factors) {
				t.Errorf("expected %v but got %v", c.expected, factors)
			}
		})
	}

}

func BenchmarkGetPrimeFactorization(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		getPrimeFactorization(64)
	}
}

func TestGetSmallestCommonMultiple(t *testing.T) {
	x := getSmallestCommonMultiple([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if expected := int64(2520); x != expected {
		t.Errorf("expected %d but got %d instead", expected, x)
	}
}

func BenchmarkGetSmallesCommonMultiple(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		getSmallestCommonMultiple([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
