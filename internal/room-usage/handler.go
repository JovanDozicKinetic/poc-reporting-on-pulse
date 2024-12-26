package roomusage

import (
	"fmt"
	"net/http"
	"path/filepath"
	helpers "pdf-poc/internal/helpers"
	"time"
)

func GenerateReportHandler(w http.ResponseWriter, r *http.Request) {

	// Default values
	dateFrom := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	dateTo := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC)
	hasIsChanged := false
	changedSince := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Parse query parameters
	// query := r.URL.Query()

	// Getting the mock data:

	dates := GetEventDataWithoutMeetingRooms()

	// Preparing the file

	fileName := fmt.Sprintf("room_usage_%s.pdf", time.Now().Format("2006-01-02-15-04-05"))

	// Preparing the template

	templateData := TemplateData{
		AreaName: "Kx Campus",
		DateFrom: dateFrom,
		DateTo:   dateTo,

		HasChangedSince: hasIsChanged,
		ChangedSince:    changedSince,

		Dates: dates,
	}

	// HTML and PDF generation

	htmlContent, err := GenerateHTML(templateData, fileName)
	if err != nil {
		http.Error(w, "Error generating HTML content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	footerText := `Room Usage and Equipment Report ` + dateFrom.Format("Mon 01 Jan 2006") + ` - ` + dateTo.Format("Mon 01 Jan 2006")

	if templateData.HasChangedSince {
		footerText += `<span style="color: #888;"> (Marked changes since: ` + templateData.ChangedSince.Format("Mon 02 Jan 2006") + `)</span>`
	}

	err = helpers.GeneratePDF(
		htmlContent,
		footerText,
		filepath.Join("pdf_exports/room-usage/", fileName),
		"templates/room-usage/header.html",
		"templates/room-usage/footer.html",
		"horizontal",
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
	http.ServeFile(w, r, filepath.Join("pdf_exports/room-usage/", fileName))
}
