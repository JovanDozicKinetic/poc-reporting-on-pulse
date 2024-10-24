package eventsrunning

import (
	"fmt"
	"net/http"
	"path/filepath"
	"pdf-poc/internal/helpers"
	"time"
)

func GenerateReportHandler(w http.ResponseWriter, r *http.Request) {

	// * No query parameters for now

	eventsRunning := GetData()

	fileName := fmt.Sprintf("events_running_%s.pdf", time.Now().Format("2006-01-02-15-04-05"))

	htmlContent, err := GenerateHTML(eventsRunning, fileName)
	if err != nil {
		http.Error(w, "Error generating HTML content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.GeneratePDF(
		htmlContent,
		"Marked changes since: "+time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC).Format("Monday 02 January 2006"),
		filepath.Join("pdf_exports/events-running/", fileName),
		"templates\\events-running\\header.html",
		"templates\\events-running\\footer.html",
		"landscape",
		helpers.Margins{
			Top:    10,
			Bottom: 10,
			Left:   10,
			Right:  10,
		})
	if err != nil {
		http.Error(w, "Error generating PDF"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filepath.Join("pdf_exports/events-running/", fileName))
}
