package easy

import (
	"encoding/json"
	"reflect"
	"testing"
)

var float64sMarshalJSONTests = []struct {
	float64s Float64s
	want     string
}{
	{
		float64s: nil,
		want:     `null`,
	},
	{
		float64s: Float64s{},
		want:     `[]`,
	},
	{
		float64s: Float64s{1.23, 0, -2.45, 1.4345531263871263},
		want:     `["1.23","0","-2.45","1.4345531263871263"]`,
	},
}

func TestFloat64s_MarshalJSON(t *testing.T) {
	for _, test := range float64sMarshalJSONTests {
		got, err := json.Marshal(test.float64s)
		if err != nil {
			t.Errorf("error occurs while json.Marshal")
			continue
		}
		if !reflect.DeepEqual(string(got), test.want) {
			t.Errorf("(%v).MarshalJSON = %v, want %v", test.float64s, string(got), test.want)
		}
	}
}

var float64sUnmarshalJSONTests = []struct {
	json string
	want Float64s
}{
	{
		json: `null`,
		want: nil,
	},
	{
		json: `[]`,
		want: Float64s{},
	},
	{
		json: `["1.23","0","-2.45","1.4345531263871263"]`,
		want: Float64s{1.23, 0, -2.45, 1.4345531263871263},
	},
	{
		json: `["1.23",0,-2.45,"1.4345531263871263"]`,
		want: Float64s{1.23, 0, -2.45, 1.4345531263871263},
	},
}

func TestFloat64s_UnmarshalJSON(t *testing.T) {
	for _, test := range float64sUnmarshalJSONTests {
		var float64s Float64s
		err := json.Unmarshal([]byte(test.json), &float64s)
		if err != nil {
			t.Errorf("error occurs while json.Marshal")
			continue
		}
		if !reflect.DeepEqual(float64s, test.want) {
			t.Errorf("(%v).UnmarshalJSON = %v, want %v", test.json, float64s, test.want)
		}
	}
}
