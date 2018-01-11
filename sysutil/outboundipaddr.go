package sysutil

import "net"

// Get Get preferred outbound ip of this machine
func GetOutboundIpaddr() string {
	conn, err := net.Dial("udp4", "1.2.3.4:56")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	if ip, _, err := net.SplitHostPort(localAddr); err != nil {
		return ""
	} else {
		return ip
	}
}
