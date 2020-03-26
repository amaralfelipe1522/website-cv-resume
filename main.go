package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
