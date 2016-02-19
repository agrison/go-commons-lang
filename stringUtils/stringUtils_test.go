package stringUtils

import "testing"

func TestAbbreviate(t *testing.T) {
	if Abbreviate("", 4) != "" {
		t.Errorf("fail test abbreviate 1")
	}
	if Abbreviate("abcdefg", 6) != "abc..." {
		t.Errorf("fail test abbreviate 2")
	}
	if Abbreviate("abcdefg", 7) != "abcdefg" {
		t.Errorf("fail test abbreviate 2")
	}
	if Abbreviate("abcdefg", 8) != "abcdefg" {
		t.Errorf("fail test abbreviate 3")
	}
	if Abbreviate("abcdefg", 4) != "a..." {
		t.Errorf("fail test abbreviate 4")
	}
}

func TestAbbreviateWithOffset(t *testing.T) {
	if AbbreviateWithOffset("", 0, 4) != "" {
		t.Errorf("fail test abbreviate 1")
	}
	if AbbreviateWithOffset("abcdefghijklmno", -1, 10) != "abcdefg..." {
		t.Errorf("fail test abbreviate 2")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 0, 10) != "abcdefg..." {
		t.Errorf("fail test abbreviate 3")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 1, 10) != "abcdefg..." {
		t.Errorf("fail test abbreviate 4")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 4, 10) != "abcdefg..." {
		t.Errorf("fail test abbreviate 5")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 5, 10) != "...fghi..." {
		t.Errorf("fail test abbreviate 5")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 6, 10) != "...ghij..." {
		t.Errorf("fail test abbreviate 6")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 8, 10) != "...ijklmno" {
		t.Errorf("fail test abbreviate 7")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 10, 10) != "...ijklmno" {
		t.Errorf("fail test abbreviate 8")
	}
	if AbbreviateWithOffset("abcdefghijklmno", 12, 10) != "...ijklmno" {
		t.Errorf("fail test abbreviate 9")
	}
	if AbbreviateWithOffset("abcdefghij", 0, 3) != "abcdefghij" {
		t.Errorf("fail test abbreviate 10")
	}
	if AbbreviateWithOffset("abcdefghij", 5, 6) != "abcdefghij" {
		t.Errorf("fail test abbreviate 11")
	}
}

func TestAppendIfMissing(t *testing.T) {
	if AppendIfMissing("abc", "") != "abc" {
		t.Errorf("fail test AppendIfMissing 1")
	}
	if AppendIfMissing("", "xyz") != "" {
		t.Errorf("fail test AppendIfMissing 2")
	}
	if AppendIfMissing("abc", "xyz") != "abcxyz" {
		t.Errorf("fail test AppendIfMissing 3")
	}
	if AppendIfMissing("abcxyz", "xyz") != "abcxyz" {
		t.Errorf("fail test AppendIfMissing 4")
	}
	if AppendIfMissing("abcXYZ", "xyz") != "abcXYZxyz" {
		t.Errorf("fail test AppendIfMissing 5")
	}
}

func TestAppendIfMissingIgnoreCase(t *testing.T) {
	if AppendIfMissingIgnoreCase("abc", "") != "abc" {
		t.Errorf("fail test AppendIfMissingIgnoreCase 1")
	}
	if AppendIfMissingIgnoreCase("", "xyz") != "" {
		t.Errorf("fail test AppendIfMissingIgnoreCase 2")
	}
	if AppendIfMissingIgnoreCase("abc", "xyz") != "abcxyz" {
		t.Errorf("fail test AppendIfMissingIgnoreCase 3")
	}
	if AppendIfMissingIgnoreCase("abcxyz", "xyz") != "abcxyz" {
		t.Errorf("fail test AppendIfMissingIgnoreCase 4")
	}
	if AppendIfMissingIgnoreCase("abcXYZ", "xyz") != "abcXYZ" {
		t.Errorf("fail test AppendIfMissingIgnoreCase 5")
	}
}

func TestCapitalize(t *testing.T) {
	if Capitalize("") != "" {
		t.Errorf("fail test capitalize 1")
	}
	if Capitalize("cat") != "Cat" {
		t.Errorf("fail test capitalize 2")
	}
	if Capitalize("cAt") != "CAt" {
		t.Errorf("fail test capitalize 3")
	}
}

func TestChomp(t *testing.T) {
	if Chomp("") != "" {
		t.Errorf("fail test chomp 1")
	}
	if Chomp("abc \r") != "abc " {
		t.Errorf("fail test chomp 2")
	}
	if Chomp("abc\n") != "abc" {
		t.Errorf("fail test chomp 3")
	}
	if Chomp("abc\r\n") != "abc" {
		t.Errorf("fail test chomp 4")
	}
	if Chomp("abc\r\n\r\n") != "abc\r\n" {
		t.Errorf("fail test chomp 5")
	}
	if Chomp("abc\n\r") != "abc\n" {
		t.Errorf("fail test chomp 6")
	}
	if Chomp("abc\n\rabc") != "abc\n\rabc" {
		t.Errorf("fail test chomp 7")
	}
	if Chomp("\r") != "" {
		t.Errorf("fail test chomp 8")
	}
	if Chomp("\n") != "" {
		t.Errorf("fail test chomp 9")
	}
	if Chomp("\r\n") != "" {
		t.Errorf("fail test chomp 10")
	}
}

func TestChop(t *testing.T) {
	if Chop("") != "" {
		t.Errorf("fail test chop 1")
	}
	if Chop("abc \r") != "abc " {
		t.Errorf("fail test chop 2")
	}
	if Chop("abc\n") != "abc" {
		t.Errorf("fail test chop 3")
	}
	if Chop("abc\r\n") != "abc" {
		t.Errorf("fail test chop 4")
	}
	if Chop("abc\r\n\r\n") != "abc\r\n" {
		t.Errorf("fail test chop 5")
	}
	if Chop("abc\n\r") != "abc\n" {
		t.Errorf("fail test chop 6")
	}
	if Chop("abc\n\rabc") != "abc\n\rab" {
		t.Errorf("fail test chop 7")
	}
	if Chop("\r") != "" {
		t.Errorf("fail test chop 8")
	}
	if Chop("\n") != "" {
		t.Errorf("fail test chop 9")
	}
	if Chop("\r\n") != "" {
		t.Errorf("fail test chop 10")
	}
}

func TestEndsWith(t *testing.T) {
	if EndsWith("foobar", "bazz") != false {
		t.Errorf("fail test EndsWith 1")
	}
	if EndsWith("foobar", "BAR") != false {
		t.Errorf("fail test EndsWith 2")
	}
	if EndsWith("foobar", "bar") != true {
		t.Errorf("fail test EndsWith 3")
	}
}

func TestEndsWithIgnoreCase(t *testing.T) {
	if EndsWithIgnoreCase("foobar", "bazz") != false {
		t.Errorf("fail test EndsWithIgnoreCase 1")
	}
	if EndsWithIgnoreCase("foobar", "BAR") != true {
		t.Errorf("fail test EndsWithIgnoreCase 2")
	}
	if EndsWithIgnoreCase("foobar", "bar") != true {
		t.Errorf("fail test EndsWithIgnoreCase 3")
	}
}

func TestIsAllLowerCase(t *testing.T) {
	if IsAllLowerCase("") != false {
		t.Errorf("fail test IsAllLowerCase 1")
	}
	if IsAllLowerCase("  ") != false {
		t.Errorf("fail test IsAllLowerCase 2")
	}
	if IsAllLowerCase("abc") != true {
		t.Errorf("fail test IsAllLowerCase 3")
	}
	if IsAllLowerCase("abC") != false {
		t.Errorf("fail test IsAllLowerCase 4")
	}
	if IsAllLowerCase("ab c") != false {
		t.Errorf("fail test IsAllLowerCase 5")
	}
	if IsAllLowerCase("ab1c") != false {
		t.Errorf("fail test IsAllLowerCase 6")
	}
	if IsAllLowerCase("ab/c") != false {
		t.Errorf("fail test IsAllLowerCase 7")
	}
}

func TestIsAllUpperCase(t *testing.T) {
	if IsAllUpperCase("") != false {
		t.Errorf("fail test IsAllUpperCase 1")
	}
	if IsAllUpperCase("  ") != false {
		t.Errorf("fail test IsAllUpperCase 2")
	}
	if IsAllUpperCase("ABC") != true {
		t.Errorf("fail test IsAllUpperCase 3")
	}
	if IsAllUpperCase("aBC") != false {
		t.Errorf("fail test IsAllUpperCase 4")
	}
	if IsAllUpperCase("A C") != false {
		t.Errorf("fail test IsAllUpperCase 5")
	}
	if IsAllUpperCase("A1C") != false {
		t.Errorf("fail test IsAllUpperCase 6")
	}
	if IsAllUpperCase("A/C") != false {
		t.Errorf("fail test IsAllUpperCase 7")
	}
}

func TestIsAlpha(t *testing.T) {
	if IsAlpha("f00b4r") != false {
		t.Errorf("fail test IsAlpha 1")
	}
	if IsAlpha("foobar") != true {
		t.Errorf("fail test IsAlpha 2")
	}
	if IsAlpha(" foo bar ") != false {
		t.Errorf("fail test IsAlpha 3")
	}
}

func TestIsAlphaSpace(t *testing.T) {
	if IsAlphaSpace("f00b4r") != false {
		t.Errorf("fail test IsAlphaSpace 1")
	}
	if IsAlphaSpace("foobar") != true {
		t.Errorf("fail test IsAlphaSpace 2")
	}
	if IsAlphaSpace(" foo bar ") != true {
		t.Errorf("fail test IsAlphaSpace 3")
	}
}

func TestIsAlphanumeric(t *testing.T) {
	if IsAlphanumeric("f00b4r") != true {
		t.Errorf("fail test IsAlphanumeric 1")
	}
	if IsAlphanumeric("foobar") != true {
		t.Errorf("fail test IsAlphanumeric 2")
	}
	if IsAlphanumeric(" foo bar ") != false {
		t.Errorf("fail test IsAlphanumeric 3")
	}
}

func TestIsAlphanumericSpace(t *testing.T) {
	if IsAlphanumericSpace("f00b4r") != true {
		t.Errorf("fail test IsAlphanumericSpace 1")
	}
	if IsAlphanumericSpace("foobar") != true {
		t.Errorf("fail test IsAlphanumericSpace 2")
	}
	if IsAlphanumericSpace(" foo4bar ") != true {
		t.Errorf("fail test IsAlphanumericSpace 3")
	}
	if IsAlphanumericSpace(". foo4bar .") != false {
		t.Errorf("fail test IsAlphanumericSpace 4")
	}
}

func TestIsNumeric(t *testing.T) {
	if IsNumeric("f00b4r") != false {
		t.Errorf("fail test IsNumeric 1")
	}
	if IsNumeric("foobar") != false {
		t.Errorf("fail test IsNumeric 2")
	}
	if IsNumeric("45667") != true {
		t.Errorf("fail test IsNumeric 3")
	}
	if IsNumeric("45 667") != false {
		t.Errorf("fail test IsNumeric 4")
	}
}

func TestIsNumericSpace(t *testing.T) {
	if IsNumericSpace("f00b4r") != false {
		t.Errorf("fail test IsNumericSpace 1")
	}
	if IsNumericSpace("foobar") != false {
		t.Errorf("fail test IsNumericSpace 2")
	}
	if IsNumericSpace("45677") != true {
		t.Errorf("fail test IsNumericSpace 3")
	}
	if IsNumericSpace("4567 123") != true {
		t.Errorf("fail test IsNumericSpace 4")
	}
	if IsNumericSpace("4567.123") != false {
		t.Errorf("fail test IsNumericSpace 5")
	}
}

func TestIsWhitespace(t *testing.T) {
	if IsWhitespace("f00b4r") != false {
		t.Errorf("fail test IsWhitespace 1")
	}
	if IsWhitespace("foo bar") != false {
		t.Errorf("fail test IsWhitespace 2")
	}
	if IsWhitespace("  ") != true {
		t.Errorf("fail test IsWhitespace 3")
	}
	if IsWhitespace("") != true {
		t.Errorf("fail test IsWhitespace 4")
	}
	if IsWhitespace("      ") != true {
		t.Errorf("fail test IsWhitespace 5")
	}
}

func TestOverlay(t *testing.T) {
	if Overlay("", "abc", 0, 0) != "abc" {
		t.Errorf("fail test Overlay 1")
	}
	if Overlay("abcdef", "", 2, 4) != "abef" {
		t.Errorf("fail test Overlay 2")
	}
	if Overlay("abcdef", "", 4, 2) != "abef" {
		t.Errorf("fail test Overlay 3")
	}
	if Overlay("abcdef", "zzzz", 2, 4) != "abzzzzef" {
		t.Errorf("fail test Overlay 4")
	}
	if Overlay("abcdef", "zzzz", 4, 2) != "abzzzzef" {
		t.Errorf("fail test Overlay 5")
	}
	if Overlay("abcdef", "zzzz", -1, 4) != "zzzzef" {
		t.Errorf("fail test Overlay 6")
	}
	if Overlay("abcdef", "zzzz", 2, 8) != "abzzzz" {
		t.Errorf("fail test Overlay 7")
	}
	if Overlay("abcdef", "zzzz", -2, -3) != "zzzzabcdef" {
		t.Errorf("fail test Overlay 8")
	}
	if Overlay("abcdef", "zzzz", 8, 10) != "abcdefzzzz" {
		t.Errorf("fail test Overlay 9")
	}
}

func TestRemove(t *testing.T) {
	if Remove("", "abc") != "" {
		t.Errorf("fail test Remove 1")
	}
	if Remove("queued", "u") != "qeed" {
		t.Errorf("fail test Remove 2")
	}
	if Remove("queued", "z") != "queued" {
		t.Errorf("fail test Remove 3")
	}
}

func TestRepeat(t *testing.T) {
	if Repeat("a", 3) != "aaa" {
		t.Errorf("fail test Repeat 1")
	}
	if Repeat("abc", 3) != "abcabcabc" {
		t.Errorf("fail test Repeat 2")
	}
}

func TestRepeatWithSeparator(t *testing.T) {
	if RepeatWithSeparator("a", ",", 3) != "a,a,a" {
		t.Errorf("fail test RepeatWithSeparator 1")
	}
	if RepeatWithSeparator("abc", "-", 3) != "abc-abc-abc" {
		t.Errorf("fail test RepeatWithSeparator 2")
	}
}

func TestStrip(t *testing.T) {
	if Strip("   abc   ") != "abc" {
		t.Errorf("fail test Strip 1")
	}
	if Strip("  . a  bc :    ") != ". a  bc :" {
		t.Errorf("fail test Strip 2")
	}
	if Strip("abc") != "abc" {
		t.Errorf("fail test Strip 3")
	}
}

func TestSubstrings(t *testing.T) {
	if SubstringBefore("abc.def.ghi", ".") != "abc" {
		t.Errorf("fail test SubstringBefore 1")
	}
	if SubstringBefore("abc.def", "g") != "abc.def" {
		t.Errorf("fail test SubstringBefore 2")
	}
	if SubstringBeforeLast("abc.def.ghi", ".") != "abc.def" {
		t.Errorf("fail test SubstringBeforeLast 1")
	}
	if SubstringAfter("abc.def.ghi", ".") != "def.ghi" {
		t.Errorf("fail test SubstringAfter 1")
	}
	if SubstringAfter("abc.def", "g") != "abc.def" {
		t.Errorf("fail test SubstringAfter 2")
	}
	if SubstringAfterLast("abc.def.ghi", ".") != "ghi" {
		t.Errorf("fail test SubstringAfterLast 1")
	}
}

func TestSwapCase(t *testing.T) {
	if SwapCase("abc.def") != "ABC.DEF" {
		t.Errorf("fail test SwapCase 1")
	}
	if SwapCase("aBc.dEf") != "AbC.DeF" {
		t.Errorf("fail test SwapCase 2")
	}
	if SwapCase(" ABC.def ") != " abc.DEF " {
		t.Errorf("fail test SwapCase 3")
	}
}

func TestUncapitalize(t *testing.T) {
	if Uncapitalize("AbcDef") != "abcDef" {
		t.Errorf("fail test Uncapitalize 1")
	}
	if Uncapitalize("abcDef") != "abcDef" {
		t.Errorf("fail test Uncapitalize 2")
	}
	if Uncapitalize("世界") != "世界" {
		t.Errorf("fail test Uncapitalize 3")
	}
}
