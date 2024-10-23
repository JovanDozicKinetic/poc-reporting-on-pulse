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
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, struct {
		EventsRunning []EventRunning
		FromDate      time.Time
		ToDate        time.Time
		// Custom parameters:
		HasIsChanged bool
	}{
		EventsRunning: eventsRunning,
		FromDate:      time.Now(),
		ToDate:        time.Now(),
		HasIsChanged:  hasIsChanged(eventsRunning), // This will be customizable from report to report
	})
	if err != nil {
		log.Fatal(err)
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
	log.Println(events[0])
	for _, event := range events {
		if event.IsChanged != nil {
			return true
		}
	}
	return false
}
