package main

import (
	"fmt"
	"net/http"
	"rest_server/api"
	"rest_server/data"
	"text/template"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request to %v ", r)
	w.Write([]byte("Hello from Go !"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) { // serverside rendering
	fmt.Printf("request to %v ", r)
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux() // factory creation of the request router or multiplexer
	port := ":3333"

	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/get", api.Get)
	server.HandleFunc("/post", api.Post)

	fmt.Printf("Server is running on %v", port)
	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("The server could not be run")
	} else {
		fmt.Printf("Server is running on %v", port)
	}
}
