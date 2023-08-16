package helpers

import (
	"testing"
)

const (
	text      string = "THIS IS A TEST STRING."
	morseCode string = "- .... .. ... / .. ... / .- / - . ... - / ... - .-. .. -. --. .-.-.-"
)

func TestStringToMorseCode(t *testing.T) {
	if StringToMorseCode(text) != morseCode {
		t.Errorf("StringToMorseCode func returned an unexpected result.")
	}
}

func TestMorseCodeToString(t *testing.T) {
	if MorseCodeToString(morseCode) != text {
		t.Errorf("MorseCodeToString func returned an unexpected result.")
	}
}
