package slice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Uniq(t *testing.T) {
	t1 := []string{}
	assert.Equal(t, t1, Uniq(t1))

	t2 := []string{"1"}
	assert.Equal(t, t2, Uniq(t2))

	t3 := []string{"1", "2"}
	for i := 0; i < 10000; i++ {
		res := Uniq(t3)
		sort.StringSlice(res).Sort()
		assert.Equal(t, t3, res)
	}

	t4 := []string{"1", "3", "5", "5", "3", "1"}
	t4_except := []string{"1", "3", "5"}
	for i := 0; i < 10000; i++ {
		res := Uniq(t4)
		sort.StringSlice(res).Sort()
		assert.Equal(t, t4_except, res)
	}
}

func Test_Union(t *testing.T) {
	s1_1 := []int{1, 2, 3, 3, 6, 6, 7, 5}
	s1_2 := []int{2, 3, 5, 1}
	s1_except := []int{1, 2, 3, 5}
	res := Union(s1_1, s1_2)
	sort.IntSlice(res).Sort()
	assert.Equal(t, s1_except, res)

	var s2_1 []int = nil
	var s2_2 []int = []int{1}
	assert.Nil(t, Union(s2_1, s2_2))

	var s3_1 []int = []int{}
	var s3_2 []int = []int{1}
	assert.Equal(t, []int{}, Union(s3_1, s3_2))
}
