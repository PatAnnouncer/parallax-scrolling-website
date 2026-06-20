package services

import (
	"net"
	"time"
)

// CheckPort checks if a port is available on the given host
func CheckPort(host string, port string) bool {
	timeout := time.Second * 3
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}
	return false
}
