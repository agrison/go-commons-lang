package wordUtils

import (
	_ "bytes"
	"regexp"
	//"strings"
	"unicode"
)

func isDelimiter(c rune, delimiters ...string) bool {
	if delimiters == nil {
		return unicode.IsSpace(c)
	}
	cs := string(c)
	for _, delimiter := range delimiters {
		if cs == delimiter {
			return true
		}
	}
	return false
}

// Capitalizes all the delimiter separated words in a String. Only the first letter of each word is changed. To convert the
// rest of each word to lowercase at the same time.
//
// The delimiters represent a set of characters understood to separate words.
// The first string character and the first non-delimiter character after a
// delimiter will be capitalized.
func Capitalize(str string) string {
	return CapitalizeDelimited(str, nil...)
}

// Capitalizes all the delimiter separated words in a String. Only the first letter of each word is changed. To convert the
// rest of each word to lowercase at the same time.
//
// The delimiters represent a set of characters understood to separate words.
// The first string character and the first non-delimiter character after a
// delimiter will be capitalized.
func CapitalizeDelimited(str string, delimiters ...string) string {
	if str == "" || (delimiters != nil && len(delimiters) == 0) {
		return str
	}
	buff := []rune(str)
	capitalizeNext := true
	for i, c := range buff {
		if isDelimiter(c, delimiters...) {
			capitalizeNext = true
		} else if capitalizeNext {
			buff[i] = unicode.ToUpper(c)
			capitalizeNext = false
		}
	}
	return string(buff[:])
}

// ContainsAllWords checks if the String contains all words.
func ContainsAllWords(str string, words ...string) bool {
	if str == "" || len(words) == 0 {
		return false
	}
	found := 0
	for _, word := range words {
		if regexp.MustCompile(`.*\b` + word + `\b.*`).MatchString(str) {
			found += 1
		}
	}
	return found == len(words)
}
