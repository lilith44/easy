package easy

import "testing"

func TestIsIPInCIDR(t *testing.T) {
	cidr := "192.168.0.1/32"
	ip := "192.168.0.1"
	if !IsIPInCIDR(ip, cidr) {
		t.Fatalf("IsIpInCIDR错误，cidr：%s， ip：%s", cidr, ip)
	}

	ip = "192.168.0"
	if IsIPInCIDR(ip, cidr) {
		t.Fatalf("IsIpInCIDR错误，cidr：%s， ip：%s", cidr, ip)
	}

	cidr = "192.168.0.1/24"
	ip = "192.168.0.15"
	if !IsIPInCIDR(ip, cidr) {
		t.Fatalf("IsIpInCIDR错误，cidr：%s， ip：%s", cidr, ip)
	}
}
