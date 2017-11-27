package util

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
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
		if isPowerOfTwo(data.i) != data.r {
			t.Errorf("expect %v for isPowerOfTwo(%v)", data.r, data.i)
		}
	}
}
