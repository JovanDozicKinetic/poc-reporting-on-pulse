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
		http.Error(w, "Missing 'sections' parameter", http.StatusBadRequest)
		return
	}

	sections, err := parseSections(sectionsParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData := SimulateEndpointCall()
	events := UnmarshalEventData(jsonData)

	siteID := 101
	fromDate, _ := time.Parse("2006-01-02", "2010-01-10")
	toDate, _ := time.Parse("2006-01-02", "2030-01-10")
	currentTime := time.Now()

	htmlContent := GenerateHTML(events, siteID, fromDate, toDate, sections)

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

func parseSections(sectionsParam string) ([]int, error) {
	var sections []int
	for _, sectionStr := range strings.Split(sectionsParam, ",") {
		section, err := strconv.Atoi(sectionStr)
		if err != nil {
			return nil, fmt.Errorf("invalid section number: %s", sectionStr)
		}
		sections = append(sections, section)
	}
	return sections, nil
}
