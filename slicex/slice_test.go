package slicex

import (
	"reflect"
	"testing"
)

type custom struct {
	Id   int64
	Name string
}

var toSliceAnyTests = []custom{
	{
		Id:   1,
		Name: "alice",
	},
	{
		Id:   2,
		Name: "bob",
	},
	{
		Id:   3,
		Name: "cindy",
	},
	{
		Id:   4,
		Name: "david",
	},
}

func TestToSliceAny(t *testing.T) {
	for i := 0; i < len(toSliceAnyTests); i++ {
		s := toSliceAnyTests[:i]
		got := ToSliceAny(s)
		for j := 0; j < len(s); j++ {
			if !reflect.DeepEqual(s[j], got[j].(custom)) {
				t.Errorf("ToSliceAny(%v) = %v", s, got)
			}
		}
	}
}

var toMapTests = []struct {
	s    []int
	want map[int]struct{}
}{
	{
		s:    nil,
		want: nil,
	},
	{
		s:    []int{},
		want: map[int]struct{}{},
	},
	{
		s: []int{1, 2, 3, 4, 5, 6},
		want: map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
			5: {},
			6: {},
		},
	},
	{
		s: []int{1, 2, 3, 3, 2, 1},
		want: map[int]struct{}{
			1: {},
			2: {},
			3: {},
		},
	},
}

func TestToMap(t *testing.T) {
	for _, test := range toMapTests {
		got := ToMap(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToMap(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var toMapFuncTests = []struct {
	s    []custom
	want map[int64]custom
}{
	{
		s:    nil,
		want: nil,
	},
	{
		s:    []custom{},
		want: map[int64]custom{},
	},
	{
		s: []custom{
			{
				Id:   1,
				Name: "alice",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   3,
				Name: "cindy",
			},
			{
				Id:   4,
				Name: "david",
			},
		},
		want: map[int64]custom{
			1: {
				Id:   1,
				Name: "alice",
			},
			2: {
				Id:   2,
				Name: "bob",
			},
			3: {
				Id:   3,
				Name: "cindy",
			},
			4: {
				Id:   4,
				Name: "david",
			},
		},
	},
	{
		s: []custom{
			{
				Id:   1,
				Name: "alice",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   1,
				Name: "alice",
			},
		},
		want: map[int64]custom{
			1: {
				Id:   1,
				Name: "alice",
			},
			2: {
				Id:   2,
				Name: "bob",
			},
		},
	},
}

func TestToMapFunc(t *testing.T) {
	for _, test := range toMapFuncTests {
		got := ToMapFunc(test.s, func(c custom) int64 {
			return c.Id
		})
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToMapFunc(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var toSliceFuncTests = []struct {
	s    []int
	want []int
	f    func(int) int
}{
	{
		s:    nil,
		want: nil,
		f: func(i int) int {
			return i
		},
	},
	{
		s:    []int{},
		want: []int{},
		f: func(i int) int {
			return i
		},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1, 4, 9, 16, 25},
		f: func(i int) int {
			return i * i
		},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{2, 3, 4, 5, 6},
		f: func(i int) int {
			return i + 1
		},
	},
}

func TestToSliceFunc(t *testing.T) {
	for _, test := range toSliceFuncTests {
		got := ToSliceFunc(test.s, test.f)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ToSliceFunc(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var deduplicateTests = []struct {
	s    []int
	want []int
}{
	{
		s:    nil,
		want: nil,
	},
	{
		s:    []int{},
		want: []int{},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		s:    []int{1, 2, 4, 4, 5},
		want: []int{1, 2, 4, 5},
	},
	{
		s:    []int{1, 2, 4, 4, 5, 23, 13, 41, 5},
		want: []int{1, 2, 4, 5, 23, 13, 41},
	},
}

func TestDeduplicate(t *testing.T) {
	for _, test := range deduplicateTests {
		got := Deduplicate(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Deduplicate(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var deduplicateFuncTests = []struct {
	s    []custom
	want []custom
}{
	{
		s:    nil,
		want: nil,
	},
	{
		s:    []custom{},
		want: []custom{},
	},
	{
		s: []custom{
			{
				Id:   1,
				Name: "alice",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   3,
				Name: "cindy",
			},
			{
				Id:   4,
				Name: "david",
			},
		},
		want: []custom{
			{
				Id:   1,
				Name: "alice",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   3,
				Name: "cindy",
			},
			{
				Id:   4,
				Name: "david",
			},
		},
	},
	{
		s: []custom{
			{
				Id:   1,
				Name: "alice",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   2,
				Name: "bob",
			},
			{
				Id:   1,
				Name: "alice",
			},
		},
		want: []custom{
			{
				Id:   1,
				Name: "alice",
			},
			{
				Id:   2,
				Name: "bob",
			},
		},
	},
}

func TestDeduplicateFunc(t *testing.T) {
	for _, test := range deduplicateFuncTests {
		got := DeduplicateFunc(test.s, func(c custom) int64 {
			return c.Id
		})
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DeduplicateFunc(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var concatTests = []struct {
	s    [][]int
	want []int
}{
	{
		s:    nil,
		want: nil,
	},
	{
		s:    [][]int{{1, 2, 3}},
		want: []int{1, 2, 3},
	},
	{
		s:    [][]int{{1, 2, 3}},
		want: []int{1, 2, 3},
	},
	{
		s:    [][]int{{4, 5, 6}, nil, {7, 8, 9}},
		want: []int{4, 5, 6, 7, 8, 9},
	},
}

func TestConcat(t *testing.T) {
	for _, test := range concatTests {
		got := Concat(test.s...)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Concat(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}
