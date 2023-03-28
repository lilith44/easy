package easy

import "strings"

var escape = map[int32]bool{
	'*':  true,
	'.':  true,
	'?':  true,
	'+':  true,
	'$':  true,
	'^':  true,
	'[':  true,
	']':  true,
	'(':  true,
	')':  true,
	'{':  true,
	'}':  true,
	'|':  true,
	'\\': true,
	'/':  true,
}

func EscapedRegexString(str string) string {
	if str == "" {
		return ""
	}

	b := strings.Builder{}
	for _, s := range str {
		if escape[s] {
			b.WriteByte('\\')
		}
		b.WriteRune(rune(s))
	}
	return b.String()
}
