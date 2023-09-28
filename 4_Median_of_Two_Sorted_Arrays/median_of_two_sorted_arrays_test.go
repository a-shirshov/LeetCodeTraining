package median

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums1 []int
	nums2 []int
	median float64
}

var casualTests = []Test {
	{[]int{1, 3}, []int{2}, 2.00000},
	//{[]int{1, 2}, []int{3, 4}, 2.50000},
}

var nullArrayTests = []Test {
	{[]int{}, []int{1}, 1},
}

func TestMedianOfTwoSortedArrays(t *testing.T) {
	for _, test := range casualTests {
		t.Run("findMedianofTwoSortedArrays", func(t *testing.T) {
		assert.Equal(t, test.median, findMedianSortedArrays(test.nums1,test.nums2))
		})
	}
}

func TestMedianOfTwoSortedArraysNull(t *testing.T) {
	for _, test := range nullArrayTests {
		t.Run("findMedianofTwoSortedArrays", func(t *testing.T) {
		assert.Equal(t, test.median, findMedianSortedArrays(test.nums1,test.nums2))
		})
	}
}