package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/reports/template1", generateReportHandler)

	log.Printf("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateReportHandler(w http.ResponseWriter, r *http.Request) {
	sectionsParam := r.URL.Query().Get("sections")
	if sectionsParam == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}

	siteIDString := r.URL.Query().Get("siteId")
	if siteIDString == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}
	siteID, err := strconv.Atoi(siteIDString)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	fromDateString := r.URL.Query().Get("from")
	if fromDateString == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}
	fromDate, err := time.Parse("2006-01-02", fromDateString)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	toDateString := r.URL.Query().Get("to")
	if toDateString == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}
	toDate, err := time.Parse("2006-01-02", toDateString)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	sections, err := parseSections(sectionsParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData := SimulateEventsEndpointCall()
	events := UnmarshalEventData(jsonData)

	jsonData = SimulateCateringTypesEndpointCall()
	cateringTypes := UnmarshalCateringTypeData(jsonData)

	currentTime := time.Now()

	htmlContent := GenerateHTML(events, cateringTypes, siteID, fromDate, toDate, sections)

	fileName := fmt.Sprintf("events_report_site_%d_%s.pdf", siteID, currentTime.Format("2006-01-02-15-04-05"))
	err = GeneratePDF(htmlContent, fileName)
	if err != nil {
		http.Error(w, "Error generating PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filepath.Join("pdf_exports/", fileName))
}

func parseSections(sectionsParam string) (string, error) {
	var sections []string
	for _, sectionStr := range strings.Split(sectionsParam, ",") {
		sectionTag := "#" + sectionStr + "#"
		sections = append(sections, sectionTag)
	}
	tags := strings.Join(sections, " ")

	log.Println("Requested sections: `", tags, "`.")
	return tags, nil
}

// func parseSections(sectionsParam string) ([]string, error) {
// 	return strings.Split(sectionsParam, ","), nil
// }
