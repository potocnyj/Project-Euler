package main

import (
	"reflect"
	"testing"
)

func TestSumInts(t *testing.T) {
	cases := []struct {
		name  string
		nums  []int
		total int
	}{
		{
			name:  "oneInt",
			nums:  []int{1},
			total: 1,
		},
		{
			name:  "exampleCase",
			nums:  []int{3, 5, 6, 9},
			total: 23,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			out := sumInts(test.nums)
			if out != test.total {
				t.Errorf("expected sum %d, got %d instead.", test.total, out)
			}
		})
	}
}

func TestGetMultiples(t *testing.T) {
	cases := []struct {
		name      string
		divs      []int
		limit     int
		multiples []int
	}{
		{
			name:      "oneInt",
			divs:      []int{1},
			limit:     2,
			multiples: []int{1},
		},
		{
			name:      "exampleCase",
			divs:      []int{3, 5},
			limit:     10,
			multiples: []int{3, 5, 6, 9},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			out := getMultiples(test.limit, test.divs...)
			if !reflect.DeepEqual(out, test.multiples) {
				t.Errorf("expected multiples %v, got %v instead.", test.multiples, out)
			}
		})
	}
}

func BenchmarkGetMultiples(b *testing.B) {
	divs := []int{3, 5}
	limit := 1000

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := getMultiples(limit, divs...)
		if len(out) == 0 {
			b.Fail()
		}
	}
}
