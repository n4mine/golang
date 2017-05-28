package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, "", reverse(""))
	assert.Equal(t, "1", reverse("1"))
	assert.Equal(t, "21", reverse("12"))
	assert.Equal(t, "321", reverse("123"))
}
