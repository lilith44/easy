package easy

import (
	"encoding/json"
	"reflect"
	"testing"
)

var int64sMarshalJSONTests = []struct {
	int64s Int64s
	want   string
}{
	{
		int64s: nil,
		want:   `null`,
	},
	{
		int64s: Int64s{},
		want:   `[]`,
	},
	{
		int64s: Int64s{1, 0, -1, 159999999999999},
		want:   `["1","0","-1","159999999999999"]`,
	},
}

func TestInt64s_MarshalJSON(t *testing.T) {
	for _, test := range int64sMarshalJSONTests {
		got, err := json.Marshal(test.int64s)
		if err != nil {
			t.Errorf("error occurs while json.Marshal")
			continue
		}
		if !reflect.DeepEqual(string(got), test.want) {
			t.Errorf("(%v).MarshalJSON = %v, want %v", test.int64s, string(got), test.want)
		}
	}
}

var int64sUnmarshalJSONTests = []struct {
	json string
	want Int64s
}{
	{
		json: `null`,
		want: nil,
	},
	{
		json: `[]`,
		want: Int64s{},
	},
	{
		json: `["1","0","-1","159999999999999"]`,
		want: Int64s{1, 0, -1, 159999999999999},
	},
	{
		json: `["1","0",-1,159999999999999]`,
		want: Int64s{1, 0, -1, 159999999999999},
	},
}

func TestInt64s_UnmarshalJSON(t *testing.T) {
	for _, test := range int64sUnmarshalJSONTests {
		var int64s Int64s
		err := json.Unmarshal([]byte(test.json), &int64s)
		if err != nil {
			t.Errorf("error occurs while json.Marshal")
			continue
		}
		if !reflect.DeepEqual(int64s, test.want) {
			t.Errorf("(%v).UnmarshalJSON = %v, want %v", test.json, int64s, test.want)
		}
	}
}
