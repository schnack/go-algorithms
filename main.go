package main

import (
	"fmt"
	"go-algorithms/merge_sort"
)

func main() {
	a := []int{9, 5, 7, 6, 8, 4, 1, 2, 3, 0}
	//insertion_sort.InsertionSortInt(a)
	merge_sort.MergeSortInt(a, 0, len(a))
	fmt.Println(a)

}
