package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/divine-within/morsetranslate-go/pkg/helpers"
	models "github.com/divine-within/morsetranslate-go/pkg/models/morsecode"
)

const (
	morsecode = "- .... .. ... / .. ... / .- / - . ... - / ... - .-. .. -. --. .-.-.-"
	text      = "THIS IS A TEST STRING."
	apiBase   = "/api/v1"
	testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.PZLMJBT9OIVG2qgp9hQr685oVYFgRgWpcSPmNcw6y7M"
)

func TestGetDictionary(t *testing.T) {
	req, err := http.NewRequest("GET", apiBase+"/dictionary", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+testToken)

	rr := httptest.NewRecorder()

	expectedResponse, err := json.MarshalIndent(helpers.Envelope{"dictionary": models.MorseCodeMap}, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(GetDictionary)
	handler.ServeHTTP(rr, req)

	if (rr.Body.String() != string(expectedResponse)) || (rr.Result().StatusCode != http.StatusOK) {
		t.Errorf("Dictionary GET request returned an unexpected result. Status: %d", rr.Result().StatusCode)
	}
}

func TestGetText(t *testing.T) {
	req, err := http.NewRequest("GET", apiBase+"/translate/morsecode?morsecode="+morsecode, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+testToken)

	rr := httptest.NewRecorder()

	expectedResponse, err := json.MarshalIndent(helpers.Envelope{"text": text}, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(GetText)
	handler.ServeHTTP(rr, req)

	if (rr.Body.String() != string(expectedResponse)) || (rr.Result().StatusCode != http.StatusOK) {
		t.Errorf("translate/morsecode GET request returned an unexpected result. Status: %d", rr.Result().StatusCode)
	}
}

func TestGetMorseCode(t *testing.T) {
	req, err := http.NewRequest("GET", apiBase+"/translate/text?text="+text, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+testToken)

	rr := httptest.NewRecorder()

	expectedResponse, err := json.MarshalIndent(helpers.Envelope{"morsecode": morsecode}, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(GetMorseCode)
	handler.ServeHTTP(rr, req)

	if (rr.Body.String() != string(expectedResponse)) || (rr.Result().StatusCode != http.StatusOK) {
		t.Errorf("translate/text GET request returned an unexpected result. Status: %d", rr.Result().StatusCode)
	}
}
