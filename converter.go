package main

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strings"
)

func ConvertImg(inputfile, extensionFrom, outputfile, extensionTo string) {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	var img image.Image
	var errDecode error

	switch extensionFrom {
	case "jpg", "jpeg":
		img, errDecode = jpeg.Decode(file)
	case "png":
		img, errDecode = png.Decode(file)
	default:
		fmt.Printf("Unknown format")
	}
	if errDecode != nil {
		log.Fatalf("Failed to decode file: %s", errDecode)
	}

	output, err := os.Create(outputfile)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer output.Close()

	switch extensionTo {
	case "jpg", "jpeg":
		err = jpeg.Encode(output, img, nil)
	case "png":
		err = png.Encode(output, img)
	}
	if err != nil {
		fmt.Printf("Failed to convert: %s", err)
	}
}

func ConvertTxt(inputfile, extensionFrom, outputfile, extensionTo string) {

	switch extensionFrom {
	case "txt":
		file, err := os.Open(inputfile)
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text() // Get the line as a string
			fmt.Println(line)
		}
	case "docx":
		content := ReadDocxFile(inputfile)

		fmt.Println(content)
	}

}

func ReadDocxFile(inputfile string) string {
	document, err := zip.OpenReader(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer document.Close()

	var docXML string
	for _, f := range document.File {
		if f.Name == "word/document.xml" {
			doc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer doc.Close()

			var builder strings.Builder
			_, _ = io.Copy(&builder, doc)
			docXML = builder.String()
			break
		}
	}

	type Text struct {
		XMLName xml.Name `xml:"t"`
		Text    string   `xml:",chardata"`
	}

	type Run struct {
		XMLName xml.Name `xml:"r"`
		Texts   []Text   `xml:"t"`
	}

	type Paragraph struct {
		XMLName xml.Name `xml:"p"`
		Runs    []Run    `xml:"r"`
	}

	type Body struct {
		XMLName    xml.Name    `xml:"body"`
		Paragraphs []Paragraph `xml:"p"`
	}

	type Document struct {
		XMLName xml.Name `xml:"document"`
		Body    Body     `xml:"body"`
	}

	var doc Document
	err = xml.Unmarshal([]byte(docXML), &doc)
	if err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

	// Print text
	for i, p := range doc.Body.Paragraphs {
		fmt.Printf("Paragraph %d: %d runs\n", i, len(p.Runs))
		for j, r := range p.Runs {
			fmt.Printf("  Run %d: %d texts\n", j, len(r.Texts))
			for _, t := range r.Texts {
				return t.Text
			}
		}
	}
	return ""
}
