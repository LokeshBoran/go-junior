package helper

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

var removeNonUTF = func(r rune) rune {
	if r == utf8.RuneError {
		return -1
	}
	return r
}

// RemoveNonUTF8Strings removes strings that isn't UTF-8 encoded
func RemoveNonUTF8Strings(string string) string {
	return strings.Map(removeNonUTF, string)
}

func (sh *StringHelper) NormalizeStr(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, err := transform.String(t, str)
	if err != nil {
		return strings.TrimSpace(str)
	}
	result = strings.ReplaceAll(result, "đ", "d")
	result = strings.ReplaceAll(result, "Đ", "D")
	// This will remove non utf characters from string
	result = RemoveNonUTF8Strings(result)
	fmt.Println(result)
	return strings.TrimSpace(result)
}

func (sh *StringHelper) HasEqualString(str1 string, str2 string, ignoreCase ...bool) bool {
	normStr1 := sh.NormalizeStr(str1)
	normStr2 := sh.NormalizeStr(str2)

	if len(ignoreCase) > 0 && !ignoreCase[0] {
		return strings.TrimSpace(normStr1) == strings.TrimSpace(normStr2)
	}
	return strings.EqualFold(strings.TrimSpace(normStr1), strings.TrimSpace(normStr2))
}
