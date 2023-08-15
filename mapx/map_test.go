package mapx

import (
	"slices"
	"sort"
	"testing"
)

type keysCase struct {
	m    map[int]int
	want []int
}

var keysTests = []keysCase{
	{
		m: map[int]int{
			1: 2,
			2: 3,
			3: 4,
		},
		want: []int{1, 2, 3},
	},
	{
		m: map[int]int{
			0:  12322412,
			3:  3,
			-4: 4,
			23: 2321312,
		},
		want: []int{-4, 0, 3, 23},
	},
	{
		m:    nil,
		want: nil,
	},
	{
		m:    map[int]int{},
		want: []int{},
	},
}

func TestKeys(t *testing.T) {
	for _, test := range keysTests {
		got := Keys(test.m)
		sort.SliceStable(got, func(i, j int) bool {
			return got[i] < got[j]
		})

		if !slices.Equal(got, test.want) {
			t.Errorf("Keys(%v) = %v, want %v", test.m, got, test.want)
		}
	}
}

type valuesCase struct {
	m    map[int]int
	want []int
}

var valuesTests = []keysCase{
	{
		m: map[int]int{
			1: 2,
			2: 3,
			3: 4,
		},
		want: []int{2, 3, 4},
	},
	{
		m: map[int]int{
			0:  12322412,
			3:  3,
			-4: 4,
			23: 2321312,
		},
		want: []int{3, 4, 2321312, 12322412},
	},
	{
		m:    nil,
		want: nil,
	},
	{
		m:    map[int]int{},
		want: []int{},
	},
}

func TestValues(t *testing.T) {
	for _, test := range valuesTests {
		got := Values(test.m)
		sort.SliceStable(got, func(i, j int) bool {
			return got[i] < got[j]
		})

		if !slices.Equal(got, test.want) {
			t.Errorf("Values(%v) = %v, want %v", test.m, got, test.want)
		}
	}
}
