package numeral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsPowerOfTwo(t *testing.T) {
	tDatas := []struct {
		i int64
		r bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
	}

	for _, data := range tDatas {
		if IsPowerOfTwo(data.i) != data.r {
			t.Errorf("expect %v for IsPowerOfTwo(%v)", data.r, data.i)
		}
	}
}

func Test_SplitByN(t *testing.T) {
	assert.Equal(t, [][2]int{[2]int{0, 51}}, SplitByN(51, 60))
	assert.Equal(t, [][2]int{[2]int{0, 51}}, SplitByN(51, 51))
	assert.Equal(t, [][2]int{[2]int{0, 50}, [2]int{50, 51}}, SplitByN(51, 50))
	assert.Equal(t, [][2]int{[2]int{0, 10}, [2]int{10, 20}, [2]int{20, 30}, [2]int{30, 40}, [2]int{40, 50}, [2]int{50, 51}}, SplitByN(51, 10))
	assert.Equal(t, [][2]int{[2]int{0, 51}}, SplitByN(51, 0))
}

func Test_SplitI(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, SplitI(123))
	assert.Equal(t, []int{0, 1}, SplitI(10))
	assert.Equal(t, []int{1}, SplitI(1))
	assert.Equal(t, []int{0}, SplitI(0))
	assert.Equal(t, []int{-1}, SplitI(-1))
	assert.Equal(t, []int{-12}, SplitI(-12))
}
