package easy

import (
	"bytes"
	"unicode"
	"unsafe"
)

// ByteToString converts bytes to string.
func ByteToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// StringToByte converts string to bytes.
func StringToByte(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// Underscore converts string to under-score case.
func Underscore(str string) string {
	if str == "" {
		return ""
	}

	buf := bytes.Buffer{}
	buf.Grow(len(str) + 1)
	for i := range str {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i != 0 {
				buf.WriteByte('_')
			}
			buf.WriteByte(str[i] - 'A' + 'a')
			continue
		}
		buf.WriteByte(str[i])
	}
	return buf.String()
}

// Camel converts string to camel case.
func Camel(str string) string {
	if str == "" {
		return ""
	}

	buf := bytes.Buffer{}
	buf.Grow(len(str))
	toUpper := false
	for i := range str {
		if str[i] == '_' {
			toUpper = true
			continue
		}
		b := str[i]
		if toUpper && b >= 'a' && b <= 'z' {
			b = b - 'a' + 'A'
		}
		buf.WriteByte(b)
		toUpper = false
	}
	return buf.String()
}

// InitialLowercase 首字母小写
func InitialLowercase(from string) (to string) {
	for i, v := range from {
		to = string(unicode.ToLower(v)) + from[i+1:]
		break
	}

	return
}
