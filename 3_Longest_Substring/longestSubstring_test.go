package longestSubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct{
	s string
	length int
} {
	{"au", 2},
	{" ", 1},
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"bwf", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		t.Run("LengthOfLongestSubstring", func(t *testing.T) {
			result := LengthOfLongestSubstring(test.s)
			assert.Equal(t,test.length,result,"Test: " + test.s)
		})
	}
}

func TestLengthOfLongestSubstringWindow(t *testing.T) {
	for _, test := range tests {
		t.Run("LengthOfLongestSubstring", func(t *testing.T) {
			result := LengthOfLongestSubstringWindow(test.s)
			assert.Equal(t,test.length,result,"Test: " + test.s)
		})
	}
}