package api

import (
	"encoding/json"
	"golang-rest-api/src/helpers"
	"golang-rest-api/src/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool              `json:"success"`
	Data    []models.Document `json:"data"`
	Errors  []string          `json:"errors"`
}

func CreateDocument(w http.ResponseWriter, req *http.Request) {
	body, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "No logra decodificar el cuerpo del servicio...", http.StatusBadRequest)
		return
	}

	var data = Data{
		Success: false,
		Data:    nil,
		Errors:  make([]string, 0),
	}

	body.Placa = strings.TrimSpace(body.Placa)
	if !helpers.IsValidPlaca(body.Placa) {
		data.Success = false
		data.Errors = append(data.Errors, "placa invalida..")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	body.Request = strings.TrimSpace(body.Request)
	if !helpers.IsValidDescription(body.Request) {
		data.Success = false
		data.Errors = append(data.Errors, "descripcion invalida..")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	documento, success := models.Insert(body.Placa, body.Request, body.Response)
	if success != true {
		data.Errors = append(data.Errors, "No se puede crear ...")
	}

	data.Success = success
	data.Data = append(data.Data, documento)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetDocuments(w http.ResponseWriter, req *http.Request) {
	var document []models.Document = models.GetAll()

	var data = Data{true, document, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todo_id := vars["id"]

	body, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "No logra decodificar el cuerpo del servicio ...", http.StatusBadRequest)
		return
	}

	var data = Data{Errors: make([]string, 0)}
	body.Request = strings.TrimSpace(body.Request)
	if !helpers.IsValidDescription(body.Request) {
		data.Success = false
		data.Errors = append(data.Errors, "descripcion invalida")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	documento, success := models.Update(todo_id, body.Request, body.Response)
	if success != true {
		data.Errors = append(data.Errors, "No se puede actualizar...")
	}

	data.Success = success
	data.Data = append(data.Data, documento)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var documento models.Document
	var success bool
	documento, success = models.Get(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, documento)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data = Data{Errors: make([]string, 0)}

	documento, success := models.Delete(id)
	if success != true {
		data.Errors = append(data.Errors, "No se puede eliminar...")
	}

	data.Success = success
	data.Data = append(data.Data, documento)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
