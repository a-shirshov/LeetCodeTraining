package happynumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappyNumber(t *testing.T) {
	assert.Equal(t, true, isHappy(19))
}