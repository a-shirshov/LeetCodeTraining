package longestpalindromicsubstring

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type test struct {
	s string
	palindrom string
}

var casualTests = []test {
	{"racecar", "racecar"},
	{"abba", "abba"},
	{"xanay","ana"},
}

func TestLongestPalindromicMedian(t *testing.T) {
	for _, test := range casualTests {
		t.Run("longestPalindromicSubstring", func(t *testing.T){
			assert.Equal(t, test.palindrom, longestPalindrome(test.s))
		})
	}
}