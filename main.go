// Package classification of Learn API
//
// Documentation for Learn API
//
// 	Schemes: http
// 	BasePath: /
// 	Version: 0.1.0
//
// swagger:meta
package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"

	"github.com/boombuler/barcode"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/boombuler/barcode/pdf417"

	"github.com/jung-kurt/gofpdf"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/tiaguinho/gosoap"
	"github.com/tyasrush/go-learn/helper/html"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// GetIPLocationResponse will hold the Soap response
type GetIPLocationResponse struct {
	GetIPLocationResult string `xml:"GetIpLocationResult"`
}

// GetIPLocationResult will
type GetIPLocationResult struct {
	XMLName xml.Name `xml:"GeoIP"`
	Country string   `xml:"Country"`
	State   string   `xml:"State"`
}

var (
	r GetIPLocationResponse
)

func Log(ctx context.Context, span ddtrace.Span, message string, args ...interface{}) {
	event := zerolog.Ctx(ctx).Debug()
	i := 0
	for _, val := range args {
		span.SetTag(message, fmt.Sprintf("%+v\n", val))
		event.Interface(strconv.Itoa(i), fmt.Sprintf("%+v\n", val))
		i++
	}

	event.Msg(fmt.Sprintf("%v . dd.trace_id=%d dd.span_id=%d ", message, span.Context().TraceID, span.Context().SpanID))
}

func main() {
	tracer.Start(tracer.WithAgentAddr("127.0.0.1:8126"))

	span, _ := tracer.StartSpanFromContext(context.Background(), "[ReservationHandler-ConnectFlight]")
	defer span.Finish()
	Log(context.Background(), span, "Testing")

	viper.SetConfigType("toml")
	viper.SetConfigFile("config/app.toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("error - %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Errorf("error not found - %v", err)
		}
	}

	cfg := viper.GetViper()

	if cfg != nil {
		psgr := cfg.Sub("postgres.read")
		fmt.Printf("config - %v\n", psgr.GetString("host"))
	}

	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient("http://wsgeoip.lavasoft.com/ipservice.asmx?WSDL", httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	// Use gosoap.ArrayParams to support fixed position params
	params := gosoap.Params{
		"sIp": "8.8.8.8",
	}

	res, err := soap.Call("GetIpLocation", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

	res.Unmarshal(&r)

	// GetIpLocationResult will be a string. We need to parse it to XML
	result := GetIPLocationResult{}
	err = xml.Unmarshal([]byte(r.GetIPLocationResult), &result)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	if result.Country != "US" {
		log.Fatalf("error: %+v", r)
	}

	log.Println("Country: ", result.Country)
	log.Println("State: ", result.State)
	str, err := os.Getwd()
	if err != nil {
		log.Fatalf("os path error: %s", err)
	}

	log.Println("path: ", str)
	log.Println("temp dir: ", os.TempDir())

	// fmt.Println("os param - ", os.Args[1])

	CreatePDF()
}

func CreatePDF() {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "bp-*.html")
	if err != nil {
		fmt.Errorf("error define temp file")
	}
	// remove after we file uploaded
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.Write([]byte(html.BoardingPassHTML)); err != nil {
		fmt.Errorf("error write file template")
	}

	// get location file from temp
	tempFileLocation := tmpFile.Name()

	// close tmpfile process
	if err = tmpFile.Close(); err != nil {
		fmt.Errorf("error close file template")
	}

	fmt.Println("file location - ", tempFileLocation)

	bytes, err := html.CreateScreenshotHTML(context.Background(), tempFileLocation)
	if err != nil {
		fmt.Errorf("error create file html")
	}

	// ioutil.ReadFile()

	// _, err = io.Copy(os.Stdout, rdr)
	err = ioutil.WriteFile("/Users/lion/Downloads/output.png", bytes, 2048)
	if err != nil {
		fmt.Errorf("error close file template")
	}

	err = os.Chmod("/Users/lion/Downloads/output.png", 0777)
	if err != nil {
		fmt.Errorf("error close file template")
	}

	// rdr, err := os.Open("/Users/lion/Downloads/output.png")
	// if err != nil {
	// 	fmt.Errorf("error close file template")
	// }

	// img, _, err := image.DecodeConfig(rdr)
	// if err != nil {
	// 	fmt.Errorf("error close file template")
	// }

	pdf := gofpdf.New("P", "mm", "Legal", "")
	pdf.AddPage()
	pdf.ImageOptions("/Users/lion/Downloads/output.png", float64(10), float64(10), 200, 0, false, gofpdf.ImageOptions{ReadDpi: false, ImageType: ""}, 0, "")

	// Set size image
	// pdf.Image(tempFileLocation, 10, 20, 190, 0, false, "", 0, "")

	// Generate file to PDF
	// var buf io.Writer
	err = pdf.OutputFileAndClose("/Users/lion/Downloads/output.pdf")
	if err != nil {
		fmt.Errorf("error close file template")
	}

	// ExampleNewPDFGenerator()

	err = CreatePDF417Barcode("testing hahaha", true)
	if err != nil {
		log.Fatal(fmt.Errorf("error %v", err))
	}
}

func ExampleNewPDFGenerator() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	rd := strings.NewReader(html.BoardingPassHTML)

	// page := wkhtmltopdf.NewPage(fmt.Sprintf("file://%s", filePath))
	page := wkhtmltopdf.NewPageReader(rd)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	// page.Zoom.Set(0.95)

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

func CreatePDF417Barcode(barcodeStr string, isRotate bool) error {
	bc, err := pdf417.Encode(barcodeStr, 5)
	if err != nil {
		return err
	}

	bcpdf, err := barcode.Scale(bc, 137, 38)
	if err != nil {
		return err
	}

	if isRotate {
		imgg := imaging.Rotate90(bcpdf)
		err = imaging.Save(imgg, "./rotate-barcode.png")
		if err != nil {
			return err
		}
	} else {
		file, err := os.Create("./barcode.png")
		if err != nil {
			return err
		}

		defer file.Close()

		png.Encode(file, bcpdf)
	}

	return nil
}
