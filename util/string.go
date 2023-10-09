package util

import (
	"bytes"
	"strings"
	"unicode"
)

const EMPTY_STRING = ""

func EmptyString(value string) bool {
	return strings.TrimSpace(value) == EMPTY_STRING
}

func AddSpace(s string) string {
	buf := &bytes.Buffer{}
	var last rune
	for i, rune := range s {
		if unicode.IsUpper(rune) && i > 0 {
			if !unicode.IsSpace(last) {
				buf.WriteRune(' ')
			}
		}
		buf.WriteRune(rune)
		last = rune
	}
	return buf.String()
}
