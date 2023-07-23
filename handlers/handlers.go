package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseGenerico struct {
	Estado  string
	Mensaje string
}

func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "GET"})
	fmt.Fprintln(response, string(output))
}

func Ejemplo_get_id(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ejemplo get_id"+vars["id"])
}

func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "ejemplo post")
}

func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ejemplo delete"+vars["id"])

}

func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ejemplo put"+vars["id"])
}
