package txtfiles

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gomutex/godocx"
)

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

func WriteDocxFile(outputfile, content string) {
	file, err := godocx.NewDocument()
	if err != nil {
		log.Fatalf("Failed to create a new file: %s", err)
	}
	defer file.Close()

	file.AddParagraph(content)

	err = file.SaveTo(outputfile)
	if err != nil {
		log.Fatalf("Failed to save file: %s", err)
	}
}
