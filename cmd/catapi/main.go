package main

import (
	"fmt"
	"github.com/rickmvi/simple-web-service/internal/domain"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./thecatapi/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/api/cat", domain.HandlerCatApi)

	fmt.Printf("Listening on port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
