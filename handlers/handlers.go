package handlers

import (
	"fmt"
	"net/http"
)

func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "ejemplo get")
}

func Ejemplo_get_id(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "ejemplo get_id")
}

func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "ejemplo post")
}

func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "ejemplo delete")
}

func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "ejemplo put")
}
