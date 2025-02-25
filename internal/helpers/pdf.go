package helpers

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Margins struct {
	Top    uint
	Bottom uint
	Left   uint
	Right  uint
}

func GeneratePDF(htmlContent, reportTitle, fileName, headerPath, footerPath, orientation string, margins Margins) error {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("Error creating new PDF generator:", err)
		return err
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(htmlContent))

	pdfg.Cover.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)
	pdfg.Title.Set(reportTitle) // we can use this as a workaround for custom footer text
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginTop.Set(margins.Top)
	pdfg.MarginBottom.Set(margins.Bottom)
	pdfg.MarginLeft.Set(margins.Left)
	pdfg.MarginRight.Set(margins.Right)
	pdfg.Dpi.Set(150)

	if orientation == "landscape" {
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	} else {
		pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	}

	if headerPath != "" {
		page.HeaderHTML.Set(headerPath)
	}
	if footerPath != "" {
		page.FooterHTML.Set(footerPath)
	}

	err = pdfg.Create()
	if err != nil {
		log.Println("Error creating a page:", err)
		return err
	}

	err = pdfg.WriteFile(fileName)
	if err != nil {
		log.Println("Error writing to the file", fileName, "error:", err)
		return err
	}

	return nil
}

func OpenFile(filename string) error {
	// Note: this only works on Windows
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", filepath.Join("pdf_exports/", filename)).Start()
	return err
}

func SaveHTMLToFile(htmlContent string, fileName string) error {

	file, err := os.Create(filepath.Join("html_exports/", fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(htmlContent)
	if err != nil {
		return err
	}

	return nil
}
