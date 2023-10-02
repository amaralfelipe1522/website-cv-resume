package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT","3000")
	port := os.Getenv("PORT")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Printf("Executando na porta %s...",port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
