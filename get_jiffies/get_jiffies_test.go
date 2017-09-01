package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getJiffies(t *testing.T) {
	assert.Equal(t, getJiffies(), 100)
}
