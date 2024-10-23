package template1

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"pdf-poc/internal/helpers"
	"strconv"
	"strings"
	"time"
)

func GenerateReportTemplate1Handler(w http.ResponseWriter, r *http.Request) {

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

	sectionsParam := r.URL.Query().Get("sections")
	if sectionsParam == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}
	sections, err := parseSections(sectionsParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, cateringTypes := GetTemplate1Data()

	fileName := fmt.Sprintf("events_report_site_%d_%s.pdf", siteID, time.Now().Format("2006-01-02-15-04-05"))

	htmlContent, err := GenerateTemplate1HTML(events, cateringTypes, siteID, fromDate, toDate, sections, fileName)
	if err != nil {
		http.Error(w, "Error generating HTML content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.GeneratePDF(
		htmlContent,
		"Events Report Title",
		filepath.Join("pdf_exports/", fileName),
		"templates\\template1\\header.html",
		"templates\\template1\\footer.html")
	if err != nil {
		http.Error(w, "Error generating PDF"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filepath.Join("pdf_exports/", fileName))
}

func parseSections(sectionsParam string) (string, error) {

	// This method parses sections part of the query (e.g. `1,2,3`) into
	//   tags that will be used to identify sections in HTML template (e.g. `#1# #2# #3#`)

	var sections []string
	for _, sectionStr := range strings.Split(sectionsParam, ",") {
		sectionTag := "#" + sectionStr + "#"
		sections = append(sections, sectionTag)
	}
	tags := strings.Join(sections, " ")

	log.Println("Requested sections: `", tags, "`.")
	return tags, nil
}
