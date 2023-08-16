package controllers

import (
	"net/http"

	"github.com/divine-within/morsetranslate-go/pkg/helpers"
	"github.com/divine-within/morsetranslate-go/pkg/models/morsecode"
)

// GET/dictionary
func GetDictionary(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"dictionary": models.MorseCodeMap})
}

// GET/text
func GetText(w http.ResponseWriter, r *http.Request) {
	morseCode := r.URL.Query().Get("morsecode")
	// TODO: Limit character length
	if morseCode == "" {
		http.Error(w, "Missing 'morsecode' parameter", http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"text": helpers.MorseCodeToString(morseCode)})
}

// GET/morsecode
func GetMorseCode(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	// TODO: Limit character length
	if text == "" {
		http.Error(w, "Missing 'text' parameter", http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"morsecode": helpers.StringToMorseCode(text)})
}
