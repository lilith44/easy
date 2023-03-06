package easy

import "testing"

func TestUnderscore(t *testing.T) {
	type Case struct {
		input  string
		expect string
	}

	cases := []*Case{
		{
			input:  "",
			expect: "",
		},
		{
			input:  "SnowFlake",
			expect: "snow_flake",
		},
		{
			input:  "snowFlake",
			expect: "snow_flake",
		},
		{
			input:  "RPC",
			expect: "r_p_c",
		},
	}

	for _, c := range cases {
		if Underscore(c.input) != c.expect {
			t.Fail()
		}
	}
}
