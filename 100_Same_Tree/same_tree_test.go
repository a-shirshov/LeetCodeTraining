package sametree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSameTree(t *testing.T) {
	var sameTreeTests = []struct {
		name string
		p *TreeNode
		q *TreeNode
		shouldBe bool
	}{
		{
			"easy false",
			&TreeNode{
				1,
				&TreeNode{2, nil, nil},
				&TreeNode{1, nil, nil},
			},
			&TreeNode{
				1,
				&TreeNode{1, nil, nil},
				&TreeNode{2, nil, nil},
			},
			false,
		},
	}
	
	for _, test := range sameTreeTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.shouldBe, isSameTree(test.p,test.q))
		})
	}
	
}