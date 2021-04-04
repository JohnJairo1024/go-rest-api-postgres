package models

import (
	"golang-rest-api/src/database"
	"log"
)

type Document struct {
	ID       int    `json:"id"`
	Placa    string `json:"placa"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

func Insert(placa string, request string, response string) (Document, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("INSERT INTO document(placa, request, response) VALUES($1, $2, $3) RETURNING id", placa, request, response).Scan(&todo_id)

	if todo_id == 0 {
		return Document{
			ID:       0,
			Placa:    "",
			Request:  "",
			Response: "",
		}, false
	}

	return Document{todo_id, placa, request, response}, true
}

func Get(placaVehiculo string) (Document, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM document WHERE placa = $1", placaVehiculo)

	var ID int
	var placa string
	var request string
	var response string
	err := row.Scan(&ID, &placa, &request, &response)
	if err != nil {
		return Document{}, false
	}

	return Document{ID, placa, request, response}, true
}

func Delete(id string) (Document, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("DELETE FROM document WHERE id = $1 RETURNING id", id).Scan(&todo_id)
	if todo_id == 0 {
		return Document{
			ID:      0,
			Placa:   "",
			Request: "",
		}, false
	}
	return Document{
		ID:      todo_id,
		Placa:   "",
		Request: "",
	}, true
}

func Update(id string, request string, response string) (Document, bool) {
	db := database.GetConnection()

	var todo_id int
	//TODO validar two resquest y response
	db.QueryRow("UPDATE document SET request = $1 WHERE id = $2 RETURNING id", request, response, id).Scan(&todo_id)
	if todo_id == 0 {
		return Document{
			ID:      0,
			Placa:   "",
			Request: "",
		}, false
	}

	return Document{todo_id, "", request, response}, true
}

func GetAll() []Document {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM document ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var document []Document
	for rows.Next() {
		t := Document{}

		var ID int
		var placa string
		var request string
		var response string

		err := rows.Scan(&ID, &placa, &request, &response)
		if err != nil {
			log.Fatal(err)
		}

		t.ID = ID
		t.Placa = placa
		t.Request = request
		t.Response = response

		document = append(document, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return document
}
