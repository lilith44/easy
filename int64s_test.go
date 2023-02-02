package easy

import (
	"reflect"
	"testing"
)

func TestInt64s_MarshalJSON(t *testing.T) {
	suit := [][]any{
		{
			Int64s{1, 2, 3, 4, 5, 6, 7},
			[]byte(`["1", "2", "3", "4", "5", "6", "7"]`),
		},
		{
			Int64s{},
			[]byte("[]"),
		},
	}

	for _, _case := range suit {
		s := _case[0].(Int64s)
		bytes, err := s.MarshalJSON()
		if err != nil {
			t.Fatal()
		}

		if !reflect.DeepEqual(bytes, _case[1]) {
			t.Fatal()
		}
	}
}

func TestInt64s_UnmarshalJSON(t *testing.T) {
	suit := [][]any{
		{
			[]byte(`["1", "2" ,  "3", "4" , "5", "6", "7"]`),
			Int64s{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]byte(`["1", 2 ,  "3", "4" , "5", "6", "7"]`),
			Int64s{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]byte(`[]`),
			Int64s{},
		},
	}

	for _, _case := range suit {
		int64s := Int64s{}
		if err := int64s.UnmarshalJSON(_case[0].([]byte)); err != nil {
			t.Fatal()
		}

		if !reflect.DeepEqual(int64s, _case[1].(Int64s)) {
			t.Fatal()
		}
	}
}

func TestInt64s_UnmarshalParam(t *testing.T) {
	suit := [][]any{
		{
			`"1", "2" ,  "3", "4" , "5", "6", "7"`,
			Int64s{1, 2, 3, 4, 5, 6, 7},
		},
		{
			`"1", 2 ,  "3", "4" , "5", "6", "7"`,
			Int64s{1, 2, 3, 4, 5, 6, 7},
		},
		{
			`1,2,3,4,5,6,7`,
			Int64s{1, 2, 3, 4, 5, 6, 7},
		},
		{
			``,
			Int64s{},
		},
	}

	for _, _case := range suit {
		int64s := Int64s{}
		if err := int64s.UnmarshalParam(_case[0].(string)); err != nil {
			t.Fatal()
		}

		if !reflect.DeepEqual(int64s, _case[1].(Int64s)) {
			t.Fatal()
		}
	}
}

func BenchmarkInt64s_MarshalJSON(b *testing.B) {
	int64s := Int64s{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		_, _ = int64s.MarshalJSON()
	}
}

func BenchmarkInt64s_UnmarshalJSON(b *testing.B) {
	int64s := make(Int64s, 0)
	data := []byte(`[1, "2", 3, 4, 5, 6, 7]`)
	for i := 0; i < b.N; i++ {
		_ = int64s.UnmarshalJSON(data)
	}
}

func BenchmarkInt64s_UnmarshalParam(b *testing.B) {
	int64s := make(Int64s, 0)
	src := "[1, 2, 3, 4, 5, 6, 7]"
	for i := 0; i < b.N; i++ {
		_ = int64s.UnmarshalParam(src)
	}
}
