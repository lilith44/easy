package easy

import "reflect"

// ToAnySlice transfer the given slice to []any
func ToAnySlice(slice any) []any {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return []any{slice}
	}

	value := reflect.ValueOf(slice)
	result := make([]any, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, value.Index(i).Interface())
	}
	return result
}
