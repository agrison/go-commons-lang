package wordUtils

import (
	_ "fmt"
	"testing"
)

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

func TestInitials(t *testing.T) {
	if Initials("") != "" {
		t.Errorf("fail test Initials 1")
	}
	if Initials("Ben John Lee") != "BJL" {
		t.Errorf("fail test Initials 2")
	}
	if Initials("Ben J.Lee") != "BJ" {
		t.Errorf("fail test Initials 3")
	}
}

func TestInitialsDelimited(t *testing.T) {
	if InitialsDelimited("", "") != "" {
		t.Errorf("fail test InitialsDelimited 1")
	}
	if InitialsDelimited("Ben J.Lee", " ", ".") != "BJL" {
		t.Errorf("fail test InitialsDelimited 2")
	}
}

func TestSwapCase(t *testing.T) {
	if SwapCase("") != "" {
		t.Errorf("fail test SwapCase 1")
	}
	if SwapCase("The dog has a BONE") != "tHE DOG HAS A bone" {
		t.Errorf("fail test SwapCase 2")
	}
}

func TestUncapitalize(t *testing.T) {
	if Uncapitalize("") != "" {
		t.Errorf("fail test Uncapitalize 1")
	}
	if Uncapitalize("I Am FINE") != "i am fINE" {
		t.Errorf("fail test Uncapitalize 2")
	}
}

func TestUncapitalizeDelimited(t *testing.T) {
	if UncapitalizeDelimited("", "") != "" {
		t.Errorf("fail test UncapitalizeDelimited 1")
	}
	if UncapitalizeDelimited("I Am Fine", " ") != "i am fine" {
		t.Errorf("fail test UncapitalizeDelimited 2")
	}
	if UncapitalizeDelimited("I AM.FINE", ".") != "i AM.fINE" {
		t.Errorf("fail test UncapitalizeDelimited 3")
	}
}
