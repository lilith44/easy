package easy

import (
	"bytes"
	"strconv"
)

// Float64s is a float64 slice but transfers to a string slice while json.Marshal.
type Float64s []float64

// MarshalJSON implements json.Marshaler.
func (f Float64s) MarshalJSON() ([]byte, error) {
	if f == nil {
		return []byte("null"), nil
	}

	buffer := bytes.Buffer{}
	buffer.WriteByte('[')
	for i := range f {
		buffer.WriteByte('"')
		buffer.WriteString(strconv.FormatFloat((f)[i], 'g', -1, 64))
		buffer.WriteByte('"')
		if i != len(f)-1 {
			buffer.Write([]byte(","))
		}
	}
	buffer.WriteByte(']')
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *Float64s) UnmarshalJSON(data []byte) error {
	return f.unmarshal(data)
}

// UnmarshalParam implements echo.BindUnmarshaler.
func (f *Float64s) UnmarshalParam(src string) error {
	return f.unmarshal(StringToByte(src))
}

func (f *Float64s) unmarshal(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	data = bytes.TrimLeft(bytes.TrimRight(data, "]"), "[")
	data = bytes.Replace(data, []byte(`"`), []byte(""), -1)
	if len(data) == 0 {
		*f = make(Float64s, 0)
		return nil
	}

	byteSlice := bytes.Split(data, []byte(","))
	float64s := make([]float64, len(byteSlice))
	for i := range byteSlice {
		byteSlice[i] = bytes.TrimSpace(byteSlice[i])

		var err error
		if float64s[i], err = strconv.ParseFloat(ByteToString(byteSlice[i]), 64); err != nil {
			return err
		}
	}

	*f = float64s
	return nil
}
