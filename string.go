package easy

import "bytes"

func Underscore(str string) string {
	if str == "" {
		return ""
	}

	var (
		buffer bytes.Buffer
		incr   uint8 = 'a' - 'A'
	)

	for i := range str {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i != 0 {
				buffer.WriteByte('_')
			}
			buffer.WriteByte(str[i] + incr)

			continue
		}
		buffer.WriteByte(str[i])
	}
	return buffer.String()
}
