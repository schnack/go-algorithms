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

	left := make([]int, countLeft)
	right := make([]int, countRight)

	for i := 0; i < countLeft; i++ {
		left[i] = data[first+i]
	}

	for i := 0; i < countRight; i++ {
		right[i] = data[second+i]
	}

	i := 0
	j := 0

	for k := first; k < count; k++ {

		if len(left) == i {
			data[k] = right[j]
			j++
		} else if len(right) == j {
			data[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
	}

}
