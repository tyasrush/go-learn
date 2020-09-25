package html

import (
	"fmt"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(templ string) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	rd := strings.NewReader(templ)

	// page := wkhtmltopdf.NewPage(fmt.Sprintf("file://%s", filePath))
	page := wkhtmltopdf.NewPageReader(rd)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	// err = pdfg.WriteFile("./simplesample.pdf")
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
