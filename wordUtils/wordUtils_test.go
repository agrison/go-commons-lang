package wordUtils

import "testing"

func TestCapitalize(t *testing.T) {
	if Capitalize("") != "" {
		t.Errorf("fail test Capitalize 1")
	}
	if Capitalize("i am FINE") != "I Am FINE" {
		t.Errorf("fail test Capitalize 2")
	}
}

func TestCapitalizeDelimited(t *testing.T) {
	if CapitalizeDelimited("", "") != "" {
		t.Errorf("fail test CapitalizeDelimited 1")
	}
	if CapitalizeDelimited("i am fine", " ") != "I Am Fine" {
		t.Errorf("fail test CapitalizeDelimited 2")
	}
	if CapitalizeDelimited("i aM.fine", ".") != "I aM.Fine" {
		t.Errorf("fail test CapitalizeDelimited 3")
	}
}

func TestContainsAllWords(t *testing.T) {
	if ContainsAllWords("", "a") != false {
		t.Errorf("fail test ContainsAllWords 1")
	}
	if ContainsAllWords("abcdf", "ab", "cd") != false {
		t.Errorf("fail test ContainsAllWords 2")
	}
	if ContainsAllWords("abc def", "def", "abc") != true {
		t.Errorf("fail test ContainsAllWords 3")
	}
}
