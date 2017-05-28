package util

import (
	"net"
	"strings"
)

// getOutboundIP Get preferred outbound ip of this machine
func getOutboundIP() string {
	conn, err := net.Dial("udp4", "1.2.3.4:56")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	ipPort := strings.Split(localAddr, ":")
	if len(ipPort) == 2 {
		return ipPort[0]
	}
	return ""
}
