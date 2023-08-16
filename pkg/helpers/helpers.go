package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	models "github.com/divine-within/morsetranslate-go/pkg/models/morsecode"
)

type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}
	return nil
}

func StringToMorseCode(s string) string {
	var morseCode []string
	for _, word := range strings.Split(s, " ") {
		var chars []string
		for _, char := range strings.Split(word, "") {
			if val, ok := models.MorseCodeMap[strings.ToUpper(char)]; ok {
				chars = append(chars, val)
			} else {
				chars = append(chars, "#")
			}
		}
		morseCode = append(morseCode, strings.Join(chars, " "))
	}

	return strings.Join(morseCode[:], " / ")
}

func MorseCodeToString(s string) string {
	var text []string
	for _, word := range strings.Split(s, " / ") {
		var chars []string
		for _, char := range strings.Split(word, " ") {
			if val, ok := models.ReverseMorseCodeMap[char]; ok {
				chars = append(chars, val)
			} else {
				chars = append(chars, "#")
			}
		}
		text = append(text, strings.Join(chars[:], ""))
	}

	return strings.Join(text[:], " ")
}
