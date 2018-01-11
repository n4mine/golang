package sysutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetJiffies(t *testing.T) {
	assert.Equal(t, GetJiffies(), 100)
}
