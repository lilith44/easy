package easy

import (
	"reflect"
	"testing"
)

var isIpInCIDRTests = []struct {
	ip   string
	cidr string
	want bool
}{
	{
		ip:   "192.168.118.256",
		cidr: "192.168.118.0/24",
		want: false,
	},
	{
		ip:   "192.168.118.62",
		cidr: "192.168.118.0/33",
		want: false,
	},
	{
		ip:   "192.168.118.62",
		cidr: "192.168.118.0/24",
		want: true,
	},
	{
		ip:   "192.168.118.15",
		cidr: "192.168.118.0/31",
		want: false,
	},
}

func TestIsIPInCIDR(t *testing.T) {
	for _, test := range isIpInCIDRTests {
		got := IsIPInCIDR(test.ip, test.cidr)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("IsIPInCIDR(%v, %v) = %v, want %v", test.ip, test.cidr, got, test.want)
		}
	}
}
