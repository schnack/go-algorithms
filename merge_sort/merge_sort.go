package merge_sort

import "math"

func MergeSortInt(data []int, first, count int) {
	if first < count-1 {
		second := int(math.Floor(float64(first+count) / float64(2)))
		// left
		MergeSortInt(data, first, second)
		// right
		MergeSortInt(data, second, count)
		MergeInt(data, first, second, count)
	}
}

func MergeInt(data []int, first, second, count int) {
	countLeft := second - first
	countRight := count - second

	left := make([]int, countLeft+1)
	right := make([]int, countRight+1)

	for i := 0; i < countLeft; i++ {
		left[i] = data[first+i]
	}

	for i := 0; i < countRight; i++ {
		right[i] = data[second+i]
	}

	left[countLeft] = int(math.MaxInt64)
	right[countRight] = int(math.MaxInt64)

	i := 0
	j := 0

	for k := first; k < count; k++ {

		if left[i] <= right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
	}

}
