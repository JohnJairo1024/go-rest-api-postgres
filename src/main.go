package main

import (
	"fmt"
	"golang-rest-api/src/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var port = "8080"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/documents", api.GetDocuments).Methods("GET")
	apiRouter.HandleFunc("/document/{placa}", api.GetDocument).Methods("GET")
	apiRouter.HandleFunc("/document", api.CreateDocument).Methods("POST")
	apiRouter.HandleFunc("/document/{id}", api.DeleteDocument).Methods("DELETE")
	apiRouter.HandleFunc("/document/{id}", api.UpdateDocument).Methods("PUT")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
