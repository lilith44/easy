package easy

import (
	"reflect"
	"testing"
)

var mergeIntervalTests = []struct {
	intervals []Interval[int, int]
	want      []Interval[int, int]
}{
	{
		intervals: []Interval[int, int]{
			{
				Left:  900,
				Right: 960,
				Power: 100,
			},
			{
				Left:  930,
				Right: 990,
				Power: 50,
			},
		},
		want: []Interval[int, int]{
			{
				Left:  900,
				Right: 930,
				Power: 100,
			},
			{
				Left:  930,
				Right: 960,
				Power: 150,
			},
			{
				Left:  960,
				Right: 990,
				Power: 50,
			},
		},
	},
	{
		intervals: []Interval[int, int]{
			{
				Left:  100,
				Right: 200,
				Power: 100,
			},
			{
				Left:  300,
				Right: 500,
				Power: 50,
			},
			{
				Left:  400,
				Right: 450,
				Power: 120,
			},
			{
				Left:  400,
				Right: 500,
				Power: 20,
			},
			{
				Left:  1000,
				Right: 1500,
				Power: 200,
			},
		},
		want: []Interval[int, int]{
			{
				Left:  100,
				Right: 200,
				Power: 100,
			},
			{
				Left:  300,
				Right: 400,
				Power: 50,
			},
			{
				Left:  400,
				Right: 450,
				Power: 190,
			},
			{
				Left:  450,
				Right: 500,
				Power: 70,
			},
			{
				Left:  1000,
				Right: 1500,
				Power: 200,
			},
		},
	},
}

func TestMergeIntervals(t *testing.T) {
	for _, test := range mergeIntervalTests {
		got := MergeIntervals(test.intervals...)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeIntervals(%v) = %v, want %v", test.intervals, got, test.want)
		}
	}
}
