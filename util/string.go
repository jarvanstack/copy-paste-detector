package util

import (
	"bytes"
	"unicode"
)

func TrimAllSpace(str string) string {
	buf := &bytes.Buffer{}
	for _, v := range str {
		if !unicode.Is(unicode.Space, v) && v != '\t' {
			buf.WriteRune(v)
		}
	}
	return buf.String()
}
