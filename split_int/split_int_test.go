package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitI(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, splitI(123))
	assert.Equal(t, []int{0, 1}, splitI(10))
	assert.Equal(t, []int{1}, splitI(1))
	assert.Equal(t, []int{0}, splitI(0))
	assert.Equal(t, []int{-1}, splitI(-1))
	assert.Equal(t, []int{-12}, splitI(-12))
}
