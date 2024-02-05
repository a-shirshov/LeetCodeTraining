package insertinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, [][]int{{1,2},{3,10},{12,16}}, insert([][]int{{1,2},{3,5},{6,7},{8,10},{12,16}},[]int{4,8}))
}