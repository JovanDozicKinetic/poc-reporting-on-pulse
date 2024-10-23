package eventsrunning

import (
	"bytes"
	"html/template"
	"log"
	helpers "pdf-poc/internal/helpers"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
)

func GenerateHTML(eventsRunning []EventRunning, fileName string) (string, error) {

	template, err := template.New("template.html").Funcs(sprig.FuncMap()).ParseFiles("templates/events-running/template.html")
	if err != nil {
		log.Println("Error while loading HTML template file:", err)
		return "", err
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, struct {
		SiteName      string
		EventsRunning []EventRunning
		FromDate      time.Time
		ToDate        time.Time
		// Custom parameters:
		HasIsChanged bool
	}{
		SiteName:      "Kx Campus",
		EventsRunning: eventsRunning,
		FromDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ToDate:        time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC),
		HasIsChanged:  hasIsChanged(eventsRunning),
	})
	if err != nil {
		log.Println("Error while generating HTML:", err)
		return "", err
	}

	htmlContent := buf.String()

	err = helpers.SaveHTMLToFile(htmlContent, strings.Replace(fileName, ".pdf", ".html", -1))
	if err != nil {
		log.Println("Error saving HTML content to a file", fileName, "error: ", err)
		return "", err
	}

	return htmlContent, nil
}

func hasIsChanged(events []EventRunning) bool {
	for _, event := range events {
		if event.IsChanged != nil {
			return true
		}
	}
	return false
}
