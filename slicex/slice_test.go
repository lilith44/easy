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

var isUniqueTests = []struct {
	s    []int
	want bool
}{
	{
		s:    nil,
		want: true,
	},
	{
		s:    []int{1},
		want: true,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: true,
	},
	{
		s:    []int{1, 2, 3, 4, 4, 5, 5},
		want: false,
	},
}

func TestIsUnique(t *testing.T) {
	for _, test := range isUniqueTests {
		got := IsUnique(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("IsUnique(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var isUniqueFuncTests = []struct {
	s      []int
	want   bool
	unique func(int) int
}{
	{
		s:    nil,
		want: true,
		unique: func(i int) int {
			return i
		},
	},
	{
		s:    []int{1},
		want: true,
		unique: func(i int) int {
			return i
		},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: true,
		unique: func(i int) int {
			return i
		},
	},
	{
		s:    []int{1, 2, 3, 4, 4, 5, 5},
		want: false,
		unique: func(i int) int {
			return i
		},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: false,
		unique: func(i int) int {
			return i % 2
		},
	},
}

func TestIsUniqueFunc(t *testing.T) {
	for _, test := range isUniqueFuncTests {
		got := IsUniqueFunc(test.s, test.unique)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("IsUniqueFunc(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var deleteFuncTests = []struct {
	s    []int
	want []int
	del  func(int) bool
}{
	{
		s:    nil,
		want: nil,
		del: func(i int) bool {
			return true
		},
	},
	{
		s:    []int{1},
		want: []int{1},
		del: func(i int) bool {
			return false
		},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1, 3, 5},
		del: func(i int) bool {
			return i%2 == 0
		},
	},
	{
		s:    []int{1, 2, 3, 4, 4, 5, 5},
		want: []int{1, 2, 4, 4, 5, 5},
		del: func(i int) bool {
			return i%3 == 0
		},
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
		del: func(i int) bool {
			return i%6 == 0
		},
	},
}

func TestDeleteFunc(t *testing.T) {
	for _, test := range deleteFuncTests {
		got := DeleteFunc(test.s, test.del)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DeleteFunc(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

var pagingTests = []struct {
	s    []int
	want []int
	page int64
	size int64
}{
	{
		s:    nil,
		want: []int{},
		page: 1,
		size: 1,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1},
		page: 1,
		size: 1,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
		page: 1,
		size: 5,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
		page: 1,
		size: 6,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{5},
		page: 5,
		size: 1,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{},
		page: 3,
		size: 3,
	},
	{
		s:    []int{1, 2, 3, 4, 5},
		want: []int{4, 5},
		page: 2,
		size: 3,
	},
}

func TestPaging(t *testing.T) {
	for _, test := range pagingTests {
		got := Paging(test.s, test.page, test.size)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Paging(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}
