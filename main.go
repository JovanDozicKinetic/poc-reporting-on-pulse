package main

import (
	"log"
	"net/http"
	eventsrunning "pdf-poc/internal/events-running"
	roomusage "pdf-poc/internal/room-usage"
	"pdf-poc/internal/template1"
)

func main() {

	http.HandleFunc("/reports/template1", template1.GenerateReportHandler)
	http.HandleFunc("/reports/events-running", eventsrunning.GenerateReportHandler)
	http.HandleFunc("/reports/room-usage", roomusage.GenerateReportHandler)

	log.Printf("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
