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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/viper"
	"github.com/tiaguinho/gosoap"
	"github.com/tyasrush/go-learn/helper/html"
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

func main() {
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
}
