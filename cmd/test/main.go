package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	response := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="pt-br">
		<head>
			<meta charset="UTF-8">
		</head>
		<body>
			<h2>Formulario enviado com sucesso!</h2>
			<p><b>Nome:</b> %s</p>
			<p><b>Endereço:</b> %s</p>
			<a href="/form.html">Voltar</a>
		</body>
		</html>`, name, address)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, response)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Minha primeira aplicacao web em Go!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Listening on port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
