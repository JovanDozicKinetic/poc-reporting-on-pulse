package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {

	jsonData := simulateEndpointCall()

	var events []EventData
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	log.Println("JSON data converted")

	siteId := 101
	fromDate, _ := time.Parse("2006-01-02", "2010-01-10")
	toDate, _ := time.Parse("2006-01-02", "2030-01-10")
	currentTime := time.Now()

	htmlContent := generateHTML(events, siteId, fromDate, toDate)

	fileName := fmt.Sprintf("events_report_site_%d_%s.pdf", siteId, currentTime.Format("2006-01-02-15-04-05"))
	err = generatePDF(htmlContent, fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("PDF report generated successfully!\nOpening %s...\n", fileName)

	err = openFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func simulateEndpointCall() []byte {
	var events []map[string]interface{}

	cateringTypes := []string{"Early Morning", "AM", "Mid Day", "PM", "Evening"}

	for i := 1; i <= 35; i++ {
		confirmed := rand.Intn(100)
		provisional := rand.Intn(50)
		event := map[string]interface{}{
			"cateringType": cateringTypes[rand.Intn(5)],
			"confirmed":    confirmed,
			"eventDate":    time.Date(2024, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z"),
			"eventId":      i,
			"provisional":  provisional,
			"total":        confirmed + provisional,
		}
		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02T15:04:05Z", events[i]["eventDate"].(string))
		dateJ, _ := time.Parse("2006-01-02T15:04:05Z", events[j]["eventDate"].(string))

		if dateI.Equal(dateJ) {
			return events[i]["cateringType"].(string) < events[j]["cateringType"].(string)
		}
		return dateI.Before(dateJ)
	})

	jsonData, err := json.Marshal(events)
	if err != nil {
		log.Fatalln("Error marshaling JSON:", err)
	}

	return jsonData
}

func generateHTML(events []EventData, siteID int, fromDate, toDate time.Time) string {
	tmpl, err := template.New("template1.html").Funcs(sprig.FuncMap()).ParseFiles("templates/template1.html")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		Events   []EventData
		SiteID   int
		FromDate time.Time
		ToDate   time.Time
		Midpoint int
	}{
		Events:   events,
		SiteID:   siteID,
		FromDate: fromDate,
		ToDate:   toDate,
		Midpoint: len(events) / 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func generatePDF(htmlContent, filename string) error {

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
		chromedp.ActionFunc(func(ctx context.Context) error {
			var buf []byte
			buf, err := page.CaptureScreenshot().Do(ctx)
			if err != nil {
				return err
			}
			return ioutil.WriteFile("debug_screenshot.png", buf, 0644)
		}),

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

func openFile(filename string) error {
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", filepath.Join("pdf_exports/", filename)).Start()
	return err
}
