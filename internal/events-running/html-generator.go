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

	convertIsChanged(templateData)

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

func convertIsChanged(templateData TemplateData) {
	if templateData.IsMeetingRoomsIncluded {
		for _, meetingRoom := range templateData.MeetingRooms {
			for i := range meetingRoom.Events {
				if meetingRoom.Events[i].IsChanged != nil && !*meetingRoom.Events[i].IsChanged {
					meetingRoom.Events[i].IsChanged = nil
				}
			}
		}
	} else {
		for i := range templateData.EventsRunning {
			if templateData.EventsRunning[i].IsChanged != nil && !*templateData.EventsRunning[i].IsChanged {
				templateData.EventsRunning[i].IsChanged = nil
			}
		}
	}
}
