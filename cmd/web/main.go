package main

import (
	"fmt"
	"github.com/rickmvi/simple-web-service/internal/domain"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/api/user", domain.Handler)

	fmt.Printf("Listening on port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
