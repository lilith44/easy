package easy

import (
	"bytes"
	"strconv"
)

// Int64s is an int64 slice but transfers to a string slice while json.Marshal.
type Int64s []int64

// MarshalJSON implements json.Marshaler.
func (i *Int64s) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteByte('[')
	for idx := range *i {
		buffer.WriteByte('"')
		buffer.WriteString(strconv.FormatInt((*i)[idx], 10))
		buffer.WriteByte('"')
		if idx != len(*i)-1 {
			buffer.Write([]byte(", "))
		}
	}
	buffer.WriteByte(']')
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int64s) UnmarshalJSON(data []byte) error {
	return i.unmarshal(data)
}

// UnmarshalParam implements echo.BindUnmarshaler.
func (i *Int64s) UnmarshalParam(src string) error {
	return i.unmarshal([]byte(src))
}

func (i *Int64s) unmarshal(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	data = bytes.TrimLeft(bytes.TrimRight(data, "]"), "[")
	data = bytes.Replace(data, []byte(`"`), []byte(""), -1)
	if len(data) == 0 {
		return nil
	}

	byteSlice := bytes.Split(data, []byte(","))
	int64s := make([]int64, len(byteSlice))
	for idx := range byteSlice {
		byteSlice[idx] = bytes.TrimSpace(byteSlice[idx])

		var err error
		if int64s[idx], err = strconv.ParseInt(ByteToString(byteSlice[idx]), 10, 64); err != nil {
			return err
		}
	}

	*i = int64s
	return nil
}
