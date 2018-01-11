package sysutil

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetOutboundIpaddr(t *testing.T) {
	ip := GetOutboundIpaddr()
	assert.NotZero(t, ip)
	assert.NotNil(t, net.ParseIP(ip))
}
