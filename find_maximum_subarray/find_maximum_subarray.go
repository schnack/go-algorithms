package find_maximum_subarray

import "math"

func FindMaxCrossingSubarray(data []int, low, mid, high int) (maxLeft, maxRight, sum int) {
	var leftSum int
	sum = 0
	for i := mid; i >= low; i-- {
		sum += data[i]
		if mid == i || sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	var rightSum int
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += data[i]
		if (mid+1) == i || sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

func FindMaxSubarray(data []int, low, high int) (maxLow, maxHigh, sum int) {
	if high == low {
		return low, high, data[low]
	} else {
		mid := int(math.Floor(float64(low+high) / float64(2)))
		leftLow, leftHigh, leftSum := FindMaxSubarray(data, low, mid)
		rightLow, rightHigh, rightSum := FindMaxSubarray(data, mid+1, high)
		crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(data, low, mid, high)

		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}

	}
}
