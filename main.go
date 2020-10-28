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
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/disintegration/imaging"

	"github.com/boombuler/barcode"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/boombuler/barcode/pdf417"

	"github.com/jung-kurt/gofpdf"
	"github.com/rs/zerolog"
	"github.com/tyasrush/go-learn/helper/html"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
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
	// tracer.Start(tracer.WithAgentAddr("127.0.0.1:8126"))

	// span, _ := tracer.StartSpanFromContext(context.Background(), "[ReservationHandler-ConnectFlight]")
	// defer span.Finish()
	// Log(context.Background(), span, "Testing")

	// viper.SetConfigType("toml")
	// viper.SetConfigFile("config/app.toml")
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Errorf("error - %v", err)
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		fmt.Errorf("error not found - %v", err)
	// 	}
	// }

	// cfg := viper.GetViper()

	// if cfg != nil {
	// 	psgr := cfg.Sub("postgres.read")
	// 	fmt.Printf("config - %v\n", psgr.GetString("host"))
	// }

	// httpClient := &http.Client{
	// 	Timeout: 1500 * time.Millisecond,
	// }
	// soap, err := gosoap.SoapClient("http://wsgeoip.lavasoft.com/ipservice.asmx?WSDL", httpClient)
	// if err != nil {
	// 	log.Fatalf("SoapClient error: %s", err)
	// }

	// Use gosoap.ArrayParams to support fixed position params
	// params := gosoap.Params{
	// 	"sIp": "8.8.8.8",
	// }

	// res, err := soap.Call("GetIpLocation", params)
	// if err != nil {
	// 	log.Fatalf("Call error: %s", err)
	// }

	// res.Unmarshal(&r)

	// GetIpLocationResult will be a string. We need to parse it to XML
	// result := GetIPLocationResult{}
	// err = xml.Unmarshal([]byte(r.GetIPLocationResult), &result)
	// if err != nil {
	// 	log.Fatalf("xml.Unmarshal error: %s", err)
	// }

	// if result.Country != "US" {
	// 	log.Fatalf("error: %+v", r)
	// }

	// log.Println("Country: ", result.Country)
	// log.Println("State: ", result.State)
	// str, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("os path error: %s", err)
	// }

	// log.Println("path: ", str)
	// log.Println("temp dir: ", os.TempDir())

	// fmt.Println("os param - ", os.Args[1])

	// CreatePDF()

	// err := CreateFlightHTMLToPDF(html.FlightEntityHTML{
	// 	PNR:        "TESTING",
	// 	BookID:     "12399999",
	// 	FlightType: "Departure",
	// 	Airline: html.FlightAirlineEntityHTML{
	// 		Logo:          "https://storage.googleapis.com/cabeen-dev/assets/icon_lion.svg",
	// 		Class:         "Economy Class",
	// 		FlightNumber:  "ID-999",
	// 		DepartureDate: "Thursday, 09 October 2020",
	// 		DepartureTime: "09:00",
	// 		ArrivedTime:   "10:00",
	// 		DurationTime:  "1h 0m",
	// 		FlightDirect:  "direct",
	// 	},
	// 	Schedules: []html.FlightScheduleEntityHTML{
	// 		html.FlightScheduleEntityHTML{
	// 			DepartureTime:            "09:00",
	// 			DepartureAirportName:     "Soekarno-Hatta International Airport",
	// 			DepartureAirportIATA:     "CGK",
	// 			DepartureAirportTerminal: "Terminal 1A",
	// 			DepartureAirportCity:     "Jakarta",
	// 			ArrivedTime:              "10:00",
	// 			ArrivedAirportName:       "Juanda International Airport",
	// 			ArrivedAirportIATA:       "SUB",
	// 			ArrivedAirportTerminal:   "Terminal 2A",
	// 			ArrivedAirportCity:       "Surabaya",
	// 			TransitCityName:          "Bandung",
	// 			TransitDurationTime:      "1h",
	// 		},
	// 		html.FlightScheduleEntityHTML{
	// 			DepartureTime:            "09:00",
	// 			DepartureAirportName:     "Soekarno-Hatta International Airport",
	// 			DepartureAirportIATA:     "CGK",
	// 			DepartureAirportTerminal: "Terminal 1A",
	// 			DepartureAirportCity:     "Jakarta",
	// 			ArrivedTime:              "10:00",
	// 			ArrivedAirportName:       "Juanda International Airport",
	// 			ArrivedAirportIATA:       "SUB",
	// 			ArrivedAirportTerminal:   "Terminal 2A",
	// 			ArrivedAirportCity:       "Surabaya",
	// 			TransitCityName:          "Bandung",
	// 			TransitDurationTime:      "1h",
	// 		},
	// 		html.FlightScheduleEntityHTML{
	// 			DepartureTime:            "09:00",
	// 			DepartureAirportName:     "Soekarno-Hatta International Airport",
	// 			DepartureAirportIATA:     "CGK",
	// 			DepartureAirportTerminal: "Terminal 1A",
	// 			DepartureAirportCity:     "Jakarta",
	// 			ArrivedTime:              "10:00",
	// 			ArrivedAirportName:       "Juanda International Airport",
	// 			ArrivedAirportIATA:       "SUB",
	// 			ArrivedAirportTerminal:   "Terminal 2A",
	// 			ArrivedAirportCity:       "Surabaya",
	// 			TransitCityName:          "Bandung",
	// 			TransitDurationTime:      "1h",
	// 		},
	// 		// html.FlightScheduleEntityHTML{
	// 		// 	DepartureTime:            "09:00",
	// 		// 	DepartureAirportName:     "Soekarno-Hatta International Airport",
	// 		// 	DepartureAirportIATA:     "CGK",
	// 		// 	DepartureAirportTerminal: "Terminal 1A",
	// 		// 	DepartureAirportCity:     "Jakarta",
	// 		// 	ArrivedTime:              "10:00",
	// 		// 	ArrivedAirportName:       "Juanda International Airport",
	// 		// 	ArrivedAirportIATA:       "SUB",
	// 		// 	ArrivedAirportTerminal:   "Terminal 2A",
	// 		// 	ArrivedAirportCity:       "Surabaya",
	// 		// 	TransitCityName:          "Bandung",
	// 		// 	TransitDurationTime:      "1h",
	// 		// },
	// 	},
	// 	Passengers: []html.PassengerEntityHTML{
	// 		html.PassengerEntityHTML{
	// 			Name:    "Paijo",
	// 			Type:    "Adult",
	// 			Baggage: 0,
	// 		},
	// 		html.PassengerEntityHTML{
	// 			Name:    "Paiji",
	// 			Type:    "Infant",
	// 			Baggage: 0,
	// 		},
	// 		html.PassengerEntityHTML{
	// 			Name:    "Paiji",
	// 			Type:    "Infant",
	// 			Baggage: 0,
	// 		},
	// 		html.PassengerEntityHTML{
	// 			Name:    "Paijo",
	// 			Type:    "Adult",
	// 			Baggage: 0,
	// 		},
	// 		html.PassengerEntityHTML{
	// 			Name:    "Paiji",
	// 			Type:    "Infant",
	// 			Baggage: 0,
	// 		},
	// 		html.PassengerEntityHTML{
	// 			Name:    "Paiji",
	// 			Type:    "Infant",
	// 			Baggage: 0,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println("error - ", err)
	// }

	// ExampleNewPDFGenerator()
	if err := CreateFlightOvernightHTMLToPDF(html.ReceiptEntityHTML{
		BookID:       "12345888",
		CreatedAt:    "Wednesday, 13 Oct 2020, 09:18",
		AdminFee:     0,
		Total:        948000,
		TotalPayment: 948000,
		Customer: struct {
			Name  string
			Email string
			Phone string
		}{
			Name:  "Putri Tanjung",
			Email: "putri@gmail.com",
			Phone: "+6278787878787",
		},
		Payment: struct {
			Method string
			Status string
		}{
			Method: "Bank Transfer",
			Status: "PAID",
		},
		Passengers: []html.PassengerEntityHTML{
			html.PassengerEntityHTML{
				Name: "Ms. Putri Tanjung",
				Type: "Adult",
			},
			html.PassengerEntityHTML{
				Name: "Ms. Fany Desriyanti",
				Type: "Adult",
			},
			html.PassengerEntityHTML{
				Name: "Ms. Fany Desriyanti",
				Type: "Adult",
			},
			html.PassengerEntityHTML{
				Name: "Ms. Fany Desriyanti",
				Type: "Adult",
			},
			html.PassengerEntityHTML{
				Name: "Ms. Fany Desriyanti",
				Type: "Adult",
			},
		},
		Items: []html.ReceiptItemEntity{
			html.ReceiptItemEntity{
				Type:       "Flight Ticket",
				Desc:       "Lion Air (Adult), CGK - SUB, Jul 20, 2020",
				Qty:        3,
				Price:      400000,
				TotalPrice: 1200000,
			},
			html.ReceiptItemEntity{
				Type:       "Flight Ticket",
				Desc:       "Lion Air (Adult), CGK - SUB, Jul 20, 2020",
				Qty:        3,
				Price:      400000,
				TotalPrice: 1200000,
			},
		},
	}); err != nil {
		fmt.Println("error - ", err)
	}
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

	pdf := gofpdf.New("P", "mm", "Legal", "")
	pdf.AddPage()
	pdf.ImageOptions("/Users/lion/Downloads/output.png", float64(10), float64(10), 200, 0, false, gofpdf.ImageOptions{ReadDpi: false, ImageType: ""}, 0, "")

	err = pdf.OutputFileAndClose("/Users/lion/Downloads/output.pdf")
	if err != nil {
		fmt.Errorf("error close file template")
	}

	ExampleNewPDFGenerator()
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
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeLegal)

	// page := wkhtmltopdf.NewPage(fmt.Sprintf("file://%s", filePath))
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html.FlightReceiptHTML))

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
	err = pdfg.WriteFile("./flight_overnight.pdf")
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

func CreateFlightHTMLToPDF(p html.FlightEntityHTML) error {
	params := []html.FlightEntityHTML{}
	switch len(p.Schedules) {
	case 1:
		if len(p.Passengers) > 14 {
			// then split into new page
			p1 := p
			p1.Index = 1
			p1.Passengers = p.Passengers[:14]
			params = append(params, p1)
			p2 := p
			p2.Index = 15
			p2.Passengers = p.Passengers[15:len(p.Passengers)]
			params = append(params, p2)
		}
		break
	case 2:
		if len(p.Passengers) > 9 {
			// then split into new page
			p1 := p
			p1.Index = 1
			p1.Passengers = p.Passengers[:9]
			params = append(params, p1)
			p2 := p
			p2.Index = 10
			p2.Passengers = p.Passengers[10:len(p.Passengers)]
			params = append(params, p2)
		}
		break
	case 3:
		if len(p.Passengers) > 3 {
			// then split into new page
			p1 := p
			p1.Index = 1
			p1.Passengers = p.Passengers[:3]
			params = append(params, p1)
			p2 := p
			p2.Index = 4
			p2.Passengers = p.Passengers[4:len(p.Passengers)]
			params = append(params, p2)
		}
		break
	}

	var flightTemplateFunc = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"add": func(x int, add int) int {
			return x + add
		},
		"now": time.Now,
	}

	bs := []*bytes.Buffer{}
	if len(params) > 1 {
		tpl, err := template.New("flight_html").Funcs(flightTemplateFunc).Parse(html.FlightHTML)
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}
		err = tpl.Execute(b, params[0])
		if err != nil {
			return err
		}

		bs = append(bs, b)

		tpl1, err := template.New("flight_second_html").Funcs(flightTemplateFunc).Parse(html.FlightSecondHTML)
		if err != nil {
			return err
		}

		b1 := &bytes.Buffer{}
		err = tpl1.Execute(b1, params[1])
		if err != nil {
			return err
		}

		bs = append(bs, b1)
	} else {
		tpl, err := template.New("flight_html").Funcs(flightTemplateFunc).Parse(html.FlightHTML)
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}
		err = tpl.Execute(b, p)
		if err != nil {
			return err
		}

		bs = append(bs, b)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	if len(bs) > 1 {

		files := []*os.File{}
		for _, el := range bs {
			if el != nil {
				tmpFile, err := ioutil.TempFile(os.TempDir(), "flight-*.html")
				if err != nil {
					return err
				}

				if _, err = tmpFile.Write(el.Bytes()); err != nil {
					return err
				}

				page := wkhtmltopdf.NewPage(fmt.Sprintf("file:///%s", tmpFile.Name()))

				// Set options for this page
				page.FooterRight.Set("[page]")
				page.FooterFontSize.Set(10)

				// Add to document
				pdfg.AddPage(page)
			}
		}

		if len(files) > 0 {
			for _, f := range files {
				if f != nil {
					if err = f.Close(); err != nil {
						return err
					}

					os.Remove(f.Name())
				}
			}
		}
	} else {
		page := wkhtmltopdf.NewPageReader(strings.NewReader(string(bs[0].Bytes())))

		// Set options for this page
		page.FooterRight.Set("[page]")
		page.FooterFontSize.Set(10)
		// page.Zoom.Set(0.95)

		// Add to document
		pdfg.AddPage(page)
	}

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./flight.pdf")
	if err != nil {
		return err
	}

	return nil
}

func CreateFlightOvernightHTMLToPDF(p html.ReceiptEntityHTML) error {
	var flightTemplateFunc = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"add": func(x int, add int) int {
			return x + add
		},
		"now": time.Now,
	}

	tpl, err := template.New("flight_receipt_html").Funcs(flightTemplateFunc).Parse(html.FlightReceiptHTML)
	if err != nil {
		return err
	}

	b := &bytes.Buffer{}
	err = tpl.Execute(b, p)
	if err != nil {
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	page := wkhtmltopdf.NewPageReader(strings.NewReader(string(b.Bytes())))

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	// page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./flight_receipt.pdf")
	if err != nil {
		return err
	}
	return nil
}
