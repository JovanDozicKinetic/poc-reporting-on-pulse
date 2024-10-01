package main

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func GenerateHTML(events []EventData, cateringTypes []CateringType, siteID int, fromDate, toDate time.Time, sections string) string {
	tmpl, err := template.New("template1.html").Funcs(sprig.FuncMap()).ParseFiles("templates/template1.html")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SECTIONS:", sections)

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

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	file, err := os.Create(filepath.Join("pdf_exports/", filename))
	if err != nil {
		return err
	}
	defer file.Close()

	encodedHTML := url.PathEscape(htmlContent)

	var pdfBuf []byte
	err = chromedp.Run(ctx, chromedp.Tasks{

		chromedp.Navigate("data:text/html," + encodedHTML),

		//This is a screenshot for debugging, use this when content does not look right
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	var buf []byte
		// 	buf, err := page.CaptureScreenshot().Do(ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	return ioutil.WriteFile("debug_screenshot.png", buf, 0644)
		// }),

		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			return err
		}),
	})
	if err != nil {
		return err
	}

	_, err = file.Write(pdfBuf)
	if err != nil {
		return err
	}

	return nil
}

// func openFile(filename string) error {
// 	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", filepath.Join("pdf_exports/", filename)).Start()
// 	return err
// }
