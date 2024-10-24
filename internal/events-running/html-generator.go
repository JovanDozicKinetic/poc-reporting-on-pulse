package eventsrunning

import (
	"bytes"
	"html/template"
	"log"
	helpers "pdf-poc/internal/helpers"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

func GenerateHTML(templateData TemplateData, fileName string) (string, error) {

	template, err := template.New("template.html").Funcs(sprig.FuncMap()).ParseFiles("templates/events-running/template.html")
	if err != nil {
		log.Println("Error while loading HTML template file:", err)
		return "", err
	}

	convertIsChanged(templateData.EventsRunning)

	var buf bytes.Buffer
	err = template.Execute(&buf, templateData)
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

func convertIsChanged(events []EventRunning) {
	for i := range events {
		if events[i].IsChanged != nil && !*events[i].IsChanged {
			events[i].IsChanged = nil
		}
	}
}
