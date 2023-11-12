package setmatrixzeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroes(t *testing.T) {
	input := [][]int{{-4,-2147483648,6,-7,0}, {-8,6,-8,-6,0}, {2147483647,2,-9,-6,-10}}
	expected := [][]int{{0,0,0,0,0},{0,0,0,0,0},{2147483647,2,-9,-6,0}}
	setZeroes(input)
	assert.Equal(t, expected, input)
}