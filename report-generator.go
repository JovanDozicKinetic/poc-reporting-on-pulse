package main

import (
	"bytes"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenerateHTML(events []EventData, cateringTypes []CateringType, siteID int, fromDate, toDate time.Time, sections string) string {
	tmpl, err := template.New("template1.html").Funcs(sprig.FuncMap()).ParseFiles("templates/template1.html")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
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
	return buf.String()
}

func GeneratePDF(htmlContent, filename string) error {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("Error creating new PDF generator:", err)
		return err
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(htmlContent))
	pdfg.AddPage(page)
	pdfg.Title.Set("Events Report Title")
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginTop.Set(20)
	pdfg.MarginBottom.Set(20)
	pdfg.MarginLeft.Set(20)
	pdfg.MarginRight.Set(20)
	pdfg.Dpi.Set(150)

	page.HeaderHTML.Set("templates\\template1_header.html")
	page.FooterHTML.Set("templates\\template1_footer.html")

	err = pdfg.Create()
	if err != nil {
		log.Println("Error creating a page:", err)
		return err
	}

	err = pdfg.WriteFile(filename)
	if err != nil {
		log.Println("Error writing to the file ", filename, ":", err)
		return err
	}

	return nil
}

// func openFile(filename string) error {
// 	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", filepath.Join("pdf_exports/", filename)).Start()
// 	return err
// }
