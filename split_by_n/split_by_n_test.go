package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitByN(t *testing.T) {
	assert.Equal(t, [][2]int{[2]int{0, 51}}, splitByN(51, 60))
	assert.Equal(t, [][2]int{[2]int{0, 51}}, splitByN(51, 51))
	assert.Equal(t, [][2]int{[2]int{0, 50}, [2]int{50, 51}}, splitByN(51, 50))
	assert.Equal(t, [][2]int{[2]int{0, 10}, [2]int{10, 20}, [2]int{20, 30}, [2]int{30, 40}, [2]int{40, 50}, [2]int{50, 51}}, splitByN(51, 10))
	assert.Equal(t, [][2]int{[2]int{0, 51}}, splitByN(51, 0))
}
