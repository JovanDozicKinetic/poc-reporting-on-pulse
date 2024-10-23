package template1

import (
	"bytes"
	"html/template"
	"log"
	helpers "pdf-poc/internal/helpers"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
)

func GenerateTemplate1HTML(events []EventData, cateringTypes []CateringType, siteID int, fromDate, toDate time.Time, sections, fileName string) (string, error) {

	template, err := template.New("template.html").Funcs(sprig.FuncMap()).ParseFiles("templates/template1/template.html")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, struct {
		Events        []EventData
		SiteID        int
		FromDate      time.Time
		ToDate        time.Time
		Sections      string
		CateringTypes []CateringType
	}{
		Events:        events,
		SiteID:        siteID,
		FromDate:      fromDate,
		ToDate:        toDate,
		Sections:      sections,
		CateringTypes: cateringTypes,
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
