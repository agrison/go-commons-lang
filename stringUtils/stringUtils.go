// Package stringUtils provides various string utilities.
package stringUtils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Abbreviate abbreviates a string using ellipses.
func Abbreviate(str string, maxWidth int) string {
	return AbbreviateWithOffset(str, 0, maxWidth)
}

// AbbreviateWithOffset abbreviates a string using ellipses at a specific offset.
func AbbreviateWithOffset(str string, offset int, maxWidth int) string {
	size := len(str)
	if str == "" || maxWidth < 4 || size <= maxWidth {
		return str
	}
	if offset > size {
		offset = size
	}
	if size-offset < maxWidth-3 {
		offset = size - (maxWidth - 3)
	}
	abbrevMarker := "..."
	if offset <= 4 {
		return str[0:maxWidth-3] + abbrevMarker
	}
	if maxWidth < 7 {
		return str
	}
	if offset+maxWidth-3 < size {
		return abbrevMarker + Abbreviate(str[offset:], maxWidth-3)
	}
	return abbrevMarker + str[size-(maxWidth-3):]
}

func internalAppendIfMissing(str string, suffix string, ignoreCase bool, suffixes ...string) string {
	if str == "" || IsEmpty(str) {
		return str
	}
	if ignoreCase {
		if EndsWithIgnoreCase(str, suffix) {
			return str
		}

		for _, suffix := range suffixes {
			if EndsWithIgnoreCase(str, (string)(suffix)) {
				return str
			}
		}
	} else {
		if EndsWith(str, suffix) {
			return str
		}

		for _, suffix := range suffixes {
			if EndsWithIgnoreCase(str, (string)(suffix)) {
				return str
			}
		}
	}

	return str + suffix
}

// AppendIfMissing appends a suffix to a string if missing.
func AppendIfMissing(str string, suffix string, suffixes ...string) string {
	return internalAppendIfMissing(str, suffix, false, suffixes...)
}

// AppendIfMissingIgnoreCase appends a suffix to a string if missing (ignoring case).
func AppendIfMissingIgnoreCase(str string, suffix string, suffixes ...string) string {
	return internalAppendIfMissing(str, suffix, true, suffixes...)
}

// Capitalize capitalizes a string changing the first letter to title case. No other letters are changed.
func Capitalize(str string) string {
	for i, c := range str {
		return string(unicode.ToUpper(c)) + str[i+1:]
	}
	return ""
}

// Chomp removes one newline from end of a string if it's there, otherwise leave it alone.
// A newline is "\n", "\r", or "\r\n".
func Chomp(str string) string {
	if str == "" {
		return str
	}
	if len(str) == 1 && (str[0:1] == "\n" || str[0:1] == "\r") {
		return ""
	}
	if EndsWith(str, "\r\n") {
		return str[:len(str)-2]
	} else if EndsWith(str, "\n") || EndsWith(str, "\r") {
		return str[:len(str)-1]
	} else {
		return str
	}
}

// Chop removes the last character from a string. If the string ends in \r\n, then remove both of them.
func Chop(str string) string {
	if str == "" {
		return ""
	}
	sc := Chomp(str)
	if len(str) > len(sc) {
		return sc
	}
	return str[0 : len(str)-1]
}

// Contains checks if string contains a search string.
func Contains(str string, search string) bool {
	return strings.Contains(str, search)
}

// ContainsAny checks if the string contains any of the string in the given array.
func ContainsAny(str string, search ...string) bool {
	for _, s := range search {
		if Contains(str, (string)(s)) {
			return true
		}
	}
	return false
}

// ContainsAnyCharacter checks if the string contains any of the character in the given string.
func ContainsAnyCharacter(str string, search string) bool {
	for _, c := range search {
		if Contains(str, (string)(c)) {
			return true
		}
	}
	return false
}

// ContainsIgnoreCase checks if the string contains the searched string ignoring case.
func ContainsIgnoreCase(str string, search string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(search))
}

// ContainsNone checks if the string contains no occurrence of searched string.
func ContainsNone(str string, search ...string) bool {
	return !ContainsAny(str, search...)
}

// ContainsNoneCharacter checks if the string contains no occurrence of searched string.
func ContainsNoneCharacter(str string, search string) bool {
	return !ContainsAnyCharacter(str, search)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ContainsOnly checks if a string contains only some strings.
func ContainsOnly(str string, search ...string) bool {
	for _, c := range str {
		if !stringInSlice((string)(c), search) {
			return false
		}
	}
	return true
}

/*func ContainsOnlyCharacter(str string, search string) bool {
	return ContainsOnly(str, strings.Split(search, ""))
}*/

// IsAllLowerCase checks if the string contains only lowercase characters.
func IsAllLowerCase(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// IsAllUpperCase checks if the string contains only uppercase characters.
func IsAllUpperCase(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

// IsAlpha checks if the string contains only Unicode letters.
func IsAlpha(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsAlphanumeric checks if the string contains only Unicode letters and digits.
func IsAlphanumeric(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsAlphaSpace checks if the string contains only Unicode letters and spaces.
func IsAlphaSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsAlphanumericSpace checks if the string contains only Unicode letters, digits and spaces.
func IsAlphanumericSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsEmpty checks if a string is empty.
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

// IsNotEmpty checks if a string is not empty.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsAnyEmpty checks if any one of the given strings are empty.
func IsAnyEmpty(strings ...string) bool {
	for _, s := range strings {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

// IsNoneEmpty checks if none of the strings are empty.
func IsNoneEmpty(strings ...string) bool {
	for _, s := range strings {
		if IsEmpty(s) {
			return false
		}
	}
	return true
}

// IsBlank checks if a string is whitespace or empty
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(s) {
		return true
	}
	return false
}

// IsNotBlank checks if a string is not empty or containing only whitespaces.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAnyBlank checks if any one of the strings are empty or containing only whitespaces.
func IsAnyBlank(strings ...string) bool {
	for _, s := range strings {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// IsNoneBlank checks if none of the strings are empty or containing only whitespaces.
func IsNoneBlank(strings ...string) bool {
	for _, s := range strings {
		if IsBlank(s) {
			return false
		}
	}
	return true
}

// IsNumeric checks if the string contains only digits.
func IsNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsNumericSpace checks if the string contains only digits and whitespace.
func IsNumericSpace(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsWhitespace checks if the string contains only whitespace.
func IsWhitespace(str string) bool {
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// Join joins an array of strings into a string where each item of the array is separated with a separator.
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

// JoinBool is the same as Join but joining boolean.
func JoinBool(a []bool, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatBool(i)
	}
	return Join(strs, sep)
}

// JoinFloat64 is the same as Join but joining float64.
// The default format given to strconv.FormatFloat is 'G' and bitSize is 32.
func JoinFloat64(a []float64, sep string) string {
	return JoinFloat64WithFormatAndPrecision(a, 'G', 32, sep)
}

// JoinFloat64WithFormatAndPrecision is the same as Join but joining float64 with a custom precision (bitSize) and format.
func JoinFloat64WithFormatAndPrecision(a []float64, fmt byte, precision int, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatFloat(i, fmt, -1, precision)
	}
	return Join(strs, sep)
}

// JoinInt is the same as Join but joining integers.
func JoinInt(a []int, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.Itoa(i)
	}
	return Join(strs, sep)
}

// JoinInt64 is the same as Join but joining int64.
func JoinInt64(a []int64, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatInt(i, 10)
	}
	return Join(strs, sep)
}

// JoinUint64 is the same as Join but joining uint64.
func JoinUint64(ints []uint64, sep string) string {
	strs := make([]string, len(ints), len(ints))
	for idx, i := range ints {
		strs[idx] = strconv.FormatUint(i, 10)
	}
	return Join(strs, sep)
}

// Left gets the leftmost len characters of a string.
func Left(str string, size int) string {
	if str == "" || size < 0 {
		return ""
	}
	if len(str) <= size {
		return str
	}
	return str[0:size]
}

/*func internalLeftPad(str string, size int, pad string) string {
	pads := size - len(str)
	if pads <= 0 {
		return str
	}
	if isEmpty(pad) {
		pad = " "
	}
	padLen := len(pad)
	strLen := len(str)
	if padLen == 1 && pads
}*/

// LowerCase converts a string to lower case.
func LowerCase(str string) string {
	return strings.ToLower(str)
}

// UpperCase converts a string to upper case.
func UpperCase(str string) string {
	return strings.ToUpper(str)
}

// Mid gets size characters from the middle of a string.
func Mid(str string, pos int, size int) string {
	if str == "" || size < 0 || pos > len(str) {
		return ""
	}
	if pos < 0 {
		pos = 0
	}
	if len(str) <= pos+size {
		return str[pos:]
	}
	return str[pos : pos+size]
}

// Overlay overlays part of a string with another string.
func Overlay(str string, overlay string, start int, end int) string {
	strLen := len(str)
	// guards
	if start < 0 {
		start = 0
	}
	if start > strLen {
		start = strLen
	}
	if end < 0 {
		end = 0
	}
	if end > strLen {
		end = strLen
	}
	if start > end {
		start, end = end, start
	}
	return str[:start] + overlay + str[end:]
}

// Remove removes all occurrences of a substring from within the source string.
func Remove(str string, remove string) string {
	if IsEmpty(str) {
		return str
	}
	index := strings.Index(str, remove)
	for index > -1 {
		str = str[:index] + str[index+len(remove):]
		index = strings.Index(str, remove)
	}
	return str
}

// RemoveEnd removes a substring only if it is at the end of a source string,
// otherwise returns the source string.
func RemoveEnd(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if EndsWith(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// RemoveEndIgnoreCase is the case insensitive removal of a substring if it is at
// the end of a source string, otherwise returns the source string.
func RemoveEndIgnoreCase(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if EndsWithIgnoreCase(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// RemovePattern removes each substring of the source string that matches
// the given regular expression
func RemovePattern(str string, pattern string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(str, "")
}

// RemoveStart removes a substring only if it is at the beginning of a source string,
// otherwise returns the source string
func RemoveStart(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if StartsWith(str, remove) {
		return str[len(remove)+1:]
	}
	return str
}

// RemoveStartIgnoreCase is the case insensitive removal of a substring if it is at
// the beginning of a source string, otherwise returns the source string.
func RemoveStartIgnoreCase(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if StartsWithIgnoreCase(str, remove) {
		return str[len(remove)+1:]
	}
	return str
}

// Repeat repeats a string `repeat` times to form a new string.
func Repeat(str string, repeat int) string {
	buff := ""
	for repeat > 0 {
		repeat = repeat - 1
		buff += str
	}
	return buff
}

// RepeatWithSeparator repeats a string `repeat` times to form a new String,
// with a string separator injected each time.
func RepeatWithSeparator(str string, sep string, repeat int) string {
	buff := ""
	for repeat > 0 {
		repeat = repeat - 1
		buff += str
		if repeat > 0 {
			buff += sep
		}
	}
	return buff
}

// Reverse reverses a string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseDelimited reverses a string separated by a delimiter.
func ReverseDelimited(str string, del string) string {
	s := strings.Split(str, del)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return Join(s, del)
}

// Right gets the rightmost len characters of a string.
func Right(str string, size int) string {
	if str == "" || size < 0 {
		return ""
	}
	if len(str) <= size {
		return str
	}
	return str[len(str)-size:]
}

// Strip strips whitespace from the start and end of a String.
func Strip(str string) string {
	return regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(str, "")
}

// StripEnd strips whitespace from the end of a String.
func StripEnd(str string) string {
	return regexp.MustCompile(`\s+$`).ReplaceAllString(str, "")
}

// StripStart strips whitespace from the start of a String.
func StripStart(str string) string {
	return regexp.MustCompile(`^\s+`).ReplaceAllString(str, "")
}

// SubstringAfter gets the substring after the first occurrence of a separator.
func SubstringAfter(str string, sep string) string {
	idx := strings.Index(str, sep)
	if idx == -1 {
		return str
	}
	return str[idx+len(sep):]
}

// SubstringAfterLast gets the substring after the last occurrence of a separator.
func SubstringAfterLast(str string, sep string) string {
	idx := strings.LastIndex(str, sep)
	if idx == -1 {
		return str
	}
	return str[idx+len(sep):]
}

// SubstringBefore gets the substring before the first occurrence of a separator.
func SubstringBefore(str string, sep string) string {
	idx := strings.Index(str, sep)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

// SubstringBeforeLast gets the substring before the last occurrence of a separator.
func SubstringBeforeLast(str string, sep string) string {
	idx := strings.LastIndex(str, sep)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

// SwapCase swaps the case of a String changing upper and title case to
// lower case, and lower case to upper case.
func SwapCase(str string) string {
	buff := ""
	for _, c := range str {
		cs := (string)(c)
		if unicode.IsLower(c) {
			buff += UpperCase(cs)
		} else if unicode.IsUpper(c) {
			buff += LowerCase(cs)
		} else {
			buff += cs
		}
	}
	return buff
}

// Trim removes control characters from both ends of this string.
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// Uncapitalize uncapitalizes a String, changing the first letter to lower case.
func Uncapitalize(str string) string {
	if str == "" {
		return str
	}
	r := []rune(str)
	if unicode.IsUpper(r[0]) {
		return string(unicode.ToLower(r[0])) + string(r[1:])
	}
	return str
}

// Wrap wraps a string with another string.
func Wrap(str string, wrapWith string) string {
	if str == "" {
		return str
	}
	return wrapWith + str + wrapWith
}

// internalStartsWith internal method to check if a string starts with a specified prefix ignoring case or not.
func internalStartsWith(str string, prefix string, ignoreCase bool) bool {
	if str == "" || prefix == "" {
		return (str == "" && prefix == "")
	}
	if utf8.RuneCountInString(prefix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix))
	}
	return strings.HasPrefix(str, prefix)
}

// StartsWith check if a string starts with a specified prefix.
func StartsWith(str string, prefix string) bool {
	return internalStartsWith(str, prefix, false)
}

// StartsWithIgnoreCase case insensitive check if a string starts with a specified prefix.
func StartsWithIgnoreCase(str string, prefix string) bool {
	return internalStartsWith(str, prefix, true)
}

// StartsWithAny check if a string starts with any of an array of specified strings.
func StartsWithAny(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if internalStartsWith(str, (string)(prefix), false) {
			return true
		}
	}
	return false
}

// StartsWithAnyIgnoreCase check if a string starts with any of an array of specified strings (ignoring case).
func StartsWithAnyIgnoreCase(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if internalStartsWith(str, (string)(prefix), true) {
			return true
		}
	}
	return false
}

// Internal method to check if a string ends with a specified suffix ignoring case or not.
func internalEndsWith(str string, suffix string, ignoreCase bool) bool {
	if str == "" || suffix == "" {
		return (str == "" && suffix == "")
	}
	if utf8.RuneCountInString(suffix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasSuffix(strings.ToLower(str), strings.ToLower(suffix))
	}
	return strings.HasSuffix(str, suffix)
}

// EndsWith check if a string ends with a specified suffix.
func EndsWith(str string, suffix string) bool {
	return internalEndsWith(str, suffix, false)
}

// EndsWithIgnoreCase case insensitive check if a string ends with a specified suffix.
func EndsWithIgnoreCase(str string, suffix string) bool {
	return internalEndsWith(str, suffix, true)
}

// EndsWithAny check if a string ends with any of an array of specified strings.
func EndsWithAny(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if internalEndsWith(str, (string)(suffix), false) {
			return true
		}
	}
	return false
}

// EndsWithAnyIgnoreCase check if a string ends with any of an array of specified strings (ignoring case).
func EndsWithAnyIgnoreCase(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if internalEndsWith(str, (string)(suffix), true) {
			return true
		}
	}
	return false
}

// Returns either the passed in String, or if the String is Empty, the value of defaultStr.
func DefaultString(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}