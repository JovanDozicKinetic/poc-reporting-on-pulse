package main

import (
	"log"
	"net/http"
	"pdf-poc/internal/template1"
)

func main() {

	http.HandleFunc("/reports/template1", template1.GenerateReportTemplate1Handler)

	log.Printf("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
