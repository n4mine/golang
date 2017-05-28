package util

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unionSlice(t *testing.T) {
	s1_1 := []int{1, 2, 3, 3, 6, 6, 7, 5}
	s1_2 := []int{2, 3, 5, 1}
	s1_except := []int{1, 2, 3, 5}
	res := unionSlice(s1_1, s1_2)
	sort.IntSlice(res).Sort()
	assert.Equal(t, s1_except, res)

	var s2_1 []int = nil
	var s2_2 []int = []int{1}
	assert.Nil(t, unionSlice(s2_1, s2_2))

	var s3_1 []int = []int{}
	var s3_2 []int = []int{1}
	assert.Equal(t, []int{}, unionSlice(s3_1, s3_2))
}
