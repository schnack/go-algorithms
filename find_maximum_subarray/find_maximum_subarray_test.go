package find_maximum_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxSubarray(t *testing.T) {
	data := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	low, high, sum := FindMaxSubarray(data, 0, 15)

	assert.Equal(t, 7, low)
	assert.Equal(t, 10, high)
	assert.Equal(t, 43, sum)
}
