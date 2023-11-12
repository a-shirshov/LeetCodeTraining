package containsduplicateii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDupl(t *testing.T) {
	assert.Equal(t, true, containsNearbyDuplicate([]int{1,0,1,1}, 1))
}