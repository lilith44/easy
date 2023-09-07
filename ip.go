package easy

import "net"

// IsIPInCIDR reports whether the ip is in the cidr.
func IsIPInCIDR(ipString string, cidrString string) bool {
	ip := net.ParseIP(ipString)
	_, cidr, err := net.ParseCIDR(cidrString)
	return err == nil && cidr.Contains(ip)
}
