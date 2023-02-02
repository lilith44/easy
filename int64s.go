package easy

import (
	"bytes"
	"strconv"
	"unsafe"
)

// Int64s is an int64 slice but transfers to a string list while json.Marshal.
type Int64s []int64

// MarshalJSON implements Marshaler of json.
func (s *Int64s) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteByte('[')
	for i := range *s {
		buffer.WriteByte('"')
		buffer.WriteString(strconv.FormatInt((*s)[i], 10))
		buffer.WriteByte('"')
		if i != len(*s)-1 {
			buffer.Write([]byte(", "))
		}
	}
	buffer.WriteByte(']')
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements Unmarshaler of json.
func (s *Int64s) UnmarshalJSON(data []byte) (err error) {
	return s.unmarshal(data)
}

// UnmarshalParam implements BindUnmarshaler of echo.
func (s *Int64s) UnmarshalParam(src string) (err error) {
	return s.unmarshal([]byte(src))
}

func (s *Int64s) unmarshal(data []byte) (err error) {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	data = bytes.TrimLeft(bytes.TrimRight(data, "]"), "[")
	data = bytes.Replace(data, []byte(`"`), []byte(""), -1)
	if len(data) == 0 {
		*s = make(Int64s, 0)
		return
	}

	byteSlice := bytes.Split(data, []byte(","))
	int64s := make([]int64, len(byteSlice))
	for i := range byteSlice {
		byteSlice[i] = bytes.TrimSpace(byteSlice[i])
		if int64s[i], err = strconv.ParseInt(*(*string)(unsafe.Pointer(&byteSlice[i])), 10, 64); err != nil {
			return err
		}
	}

	*s = int64s
	return
}
