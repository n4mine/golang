package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitN(t *testing.T) {
	assert.Equal(t, [][]int{[]int{0, 51}}, splitN(51, 60))
	assert.Equal(t, [][]int{[]int{0, 51}}, splitN(51, 51))
	assert.Equal(t, [][]int{[]int{0, 50}, []int{50, 51}}, splitN(51, 50))
	assert.Equal(t, [][]int{[]int{0, 10}, []int{10, 20}, []int{20, 30}, []int{30, 40}, []int{40, 50}, []int{50, 51}}, splitN(51, 10))
	assert.Equal(t, [][]int{[]int{0, 51}}, splitN(51, 0))
}
