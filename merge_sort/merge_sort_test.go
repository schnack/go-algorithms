package merge_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	a := []int{9, 5, 7, 6, 8, 4, 1, 2, 3, 0}
	MergeSortInt(a, 0, len(a))

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, a)
}
