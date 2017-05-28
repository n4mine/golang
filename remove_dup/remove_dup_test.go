package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDup(t *testing.T) {
	t1 := []string{}
	assert.Equal(t, t1, removeDup(t1))

	t2 := []string{"1"}
	assert.Equal(t, t2, removeDup(t2))

	t3 := []string{"1", "2"}
	for i := 0; i < 10000; i++ {
		res := removeDup(t3)
		sort.StringSlice(res).Sort()
		assert.Equal(t, t3, res)
	}

	t4 := []string{"1", "3", "5", "5", "3", "1"}
	t4_except := []string{"1", "3", "5"}
	for i := 0; i < 10000; i++ {
		res := removeDup(t4)
		sort.StringSlice(res).Sort()
		assert.Equal(t, t4_except, res)
	}
}
