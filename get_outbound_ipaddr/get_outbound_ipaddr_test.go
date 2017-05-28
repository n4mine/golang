package util

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getOutboundIP(t *testing.T) {
	ip := getOutboundIP()
	assert.NotZero(t, ip)
	assert.NotNil(t, net.ParseIP(ip))
}
