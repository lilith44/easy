package easy

import "unsafe"

// ByteToString converts bytes to string.
func ByteToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// StringToByte converts string to bytes.
func StringToByte(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
