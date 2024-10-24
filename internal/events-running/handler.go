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

	dateFrom := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	dateTo := time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC)

	hasIsChanged := true
	changedSince := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	eventsRunning := GetEventDataWithoutMeetingRooms(hasIsChanged)

	fileName := fmt.Sprintf("events_running_%s.pdf", time.Now().Format("2006-01-02-15-04-05"))

	templateData := TemplateData{
		SiteName:               "Kx Campus", // ? Will we get this from an endpoint call?
		DateFrom:               dateFrom,
		DateTo:                 dateTo,
		HasChangedSince:        hasIsChanged,
		ChangedSince:           changedSince,
		IsMeetingRoomsIncluded: false,
		EventsRunning:          eventsRunning,
		MeetingRooms:           nil,
	}

	htmlContent, err := GenerateHTML(templateData, fileName)
	if err != nil {
		http.Error(w, "Error generating HTML content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var footerText string
	if templateData.HasChangedSince {
		footerText = "Marked changes since: " + templateData.ChangedSince.Format("Monday 02 January 2006")
	} else {
		footerText = ""
	}

	err = helpers.GeneratePDF(
		htmlContent,
		footerText,
		filepath.Join("pdf_exports/events-running/", fileName),
		"templates/events-running/header.html",
		"templates/events-running/footer.html",
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
