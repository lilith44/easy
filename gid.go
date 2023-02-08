package easy

import (
	"bytes"
	"runtime"
	"strconv"
)

// Gid returns the current goroutine id.
func Gid() uint64 {
	buffer := make([]byte, 64)
	buffer = buffer[:runtime.Stack(buffer, false)]

	start := len("goroutine ")
	gid, _ := strconv.ParseUint(string(buffer[start:bytes.Index(buffer[start:], []byte(" "))+start]), 10, 64)
	return gid
}
