package main

import (
	"api-rest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()
	prefijo := "/api/v1/"
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.Ejemplo_get_id).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.Ejemplo_put).Methods("PUT")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.Ejemplo_delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", mux))

}
