// Package wordUtils provides some string utilities regarding words.
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

// Capitalizes all the whitespace separated words in a String. Only the first letter of each word is changed.
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

// Extracts the initial letters from each word in the String.
//
// The first letter of the string and all first letters after the
// defined delimiters are returned as a new string.
// Their case is not changed.
func Initials(str string) string {
	return InitialsDelimited(str, nil...)
}

// Extracts the initial letters from each word in the String.
//
// The first letter of the string and all first letters after the
// defined delimiters are returned as a new string.
// Their case is not changed.
func InitialsDelimited(str string, delimiters ...string) string {
	if str == "" || (delimiters != nil && len(delimiters) == 0) {
		return str
	}
	strLen := len(str)
	buff := make([]rune, strLen/2+1)
	count := 0
	lastWasGap := true
	for _, ch := range str {
		if isDelimiter(ch, delimiters...) {
			lastWasGap = true
		} else if lastWasGap {
			buff[count] = ch
			count += 1
			lastWasGap = false
		}
	}
	return string(buff[:count])
}

// Swaps the case of a String using a word based algorithm.
//
// Upper case character converts to Lower case
//
// Title case character converts to Lower case
//
// Lower case character after Whitespace or at start converts to Title case
//
// Other Lower case character converts to Upper case
func SwapCase(str string) string {
	if str == "" {
		return str
	}
	buff := []rune(str)
	whitespace := true
	for i, ch := range buff {
		if unicode.IsUpper(ch) || unicode.IsTitle(ch) {
			buff[i] = unicode.ToLower(ch)
			whitespace = false
		} else if unicode.IsLower(ch) {
			if whitespace {
				buff[i] = unicode.ToTitle(ch)
				whitespace = false
			} else {
				buff[i] = unicode.ToUpper(ch)
			}
		} else {
			whitespace = unicode.IsSpace(ch)
		}
	}
	return string(buff)
}

// Uncapitalizes all the whitespace separated words in a string. Only the first letter of each word is changed.
func Uncapitalize(str string) string {
	return UncapitalizeDelimited(str, nil...)
}

// Uncapitalizes all the whitespace separated words in a String. Only the first letter of each word is changed.
//
// The delimiters represent a set of characters understood to separate words.
// The first string character and the first non-delimiter character after a delimiter will be uncapitalized.
func UncapitalizeDelimited(str string, delimiters ...string) string {
	if str == "" || (delimiters != nil && len(delimiters) == 0) {
		return str
	}
	buff := []rune(str)
	uncapitalizeNext := true
	for i, c := range buff {
		if isDelimiter(c, delimiters...) {
			uncapitalizeNext = true
		} else if uncapitalizeNext {
			buff[i] = unicode.ToLower(c)
			uncapitalizeNext = false
		}
	}
	return string(buff[:])
}
