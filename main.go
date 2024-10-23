package main

import (
	"log"
	"net/http"
	eventsrunning "pdf-poc/internal/events-running"
	"pdf-poc/internal/template1"
)

func main() {

	http.HandleFunc("/reports/template1", template1.GenerateReportHandler)
	http.HandleFunc("/reports/events-running", eventsrunning.GenerateReportHandler)

	log.Printf("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
