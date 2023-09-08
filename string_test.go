package easy

import (
	"testing"
)

type custom struct {
	input string
	want  string
}

var underscoreTests = []*custom{
	{
		input: "",
		want:  "",
	},
	{
		input: "black_magician",
		want:  "black_magician",
	},
	{
		input: "blackMagicianGirl",
		want:  "black_magician_girl",
	},
	{
		input: "BlueEyesChaosMaxDragon",
		want:  "blue_eyes_chaos_max_dragon",
	},
	{
		input: "appV1.1.0",
		want:  "app_v1.1.0",
	},
	{
		input: "BlackMagician",
		want:  "black_magician",
	},
}

func TestUnderscore(t *testing.T) {
	for _, test := range underscoreTests {
		got := Underscore(test.input)
		if got != test.want {
			t.Errorf("Underscore(%s) = %s, want %s", test.input, got, test.want)
		}
	}
}

var camelTests = []*custom{
	{
		input: "",
		want:  "",
	},
	{
		input: "black_magician",
		want:  "blackMagician",
	},
	{
		input: "black_magician_girl",
		want:  "blackMagicianGirl",
	},
	{
		input: "blue_eyes_chaos_max_dragon",
		want:  "blueEyesChaosMaxDragon",
	},
	{
		input: "app_v1.1.0",
		want:  "appV1.1.0",
	},
}

func TestCamel(t *testing.T) {
	for _, test := range camelTests {
		got := Camel(test.input)
		if got != test.want {
			t.Errorf("Camel(%s) = %s, want %s", test.input, got, test.want)
		}
	}
}
