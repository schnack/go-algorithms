package main

import (
	"go-algorithms/insertion_sort"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	var a []int
	for i := 0; i < b.N; i++ {
		a = append(a, i)
		insertion_sort.InsertionSortInt(a)
	}
}
