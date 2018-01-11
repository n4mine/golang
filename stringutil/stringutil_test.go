package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	assert.Equal(t, "", Reverse(""))
	assert.Equal(t, "1", Reverse("1"))
	assert.Equal(t, "21", Reverse("12"))
	assert.Equal(t, "321", Reverse("123"))
}
