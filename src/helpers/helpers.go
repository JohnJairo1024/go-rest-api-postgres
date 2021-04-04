package helpers

import (
	"encoding/json"
	"golang-rest-api/src/models"
	"net/http"
	"strings"
)

func DecodeBody(req *http.Request) (models.Document, bool) {
	var documento models.Document
	err := json.NewDecoder(req.Body).Decode(&documento)
	if err != nil {
		return models.Document{
			ID:      0,
			Placa:   "",
			Request: "",
		}, false
	}

	return documento, true
}

func IsValidPlaca(placa string) bool {
	plate := strings.TrimSpace(placa)
	if len(plate) == 0 {
		return false
	}

	return true
}

func IsValidDescription(request string) bool {
	desc := strings.TrimSpace(request)
	if len(desc) == 0 {
		return false
	}

	return true
}
