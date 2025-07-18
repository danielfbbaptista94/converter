package txtfiles

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
)

type OfficeDoc struct {
	XMLName xml.Name   `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 document-content"`
	Body    OfficeBody `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 body"`
}

type OfficeBody struct {
	Text OfficeText `xml:"urn:oasis:names:tc:opendocument:xmlns:office:1.0 text"`
}

type OfficeText struct {
	Paragraphs []TextP `xml:"urn:oasis:names:tc:opendocument:xmlns:text:1.0 p"`
}

type TextP struct {
	Content string `xml:",chardata"`
}

func ReadOdtFile(inputfile string) string {
	file, err := zip.OpenReader(inputfile)
	if err != nil {
		fmt.Printf("Error opening ODT file: %s", inputfile)
	}
	defer file.Close()

	for _, doc := range file.File {
		if doc.Name == "content.xml" {
			docXml, err := doc.Open()
			if err != nil {
				fmt.Printf("Error opening content.xml: %s", err)
			}
			defer docXml.Close()

			contentBytes, err := io.ReadAll(docXml)
			if err != nil {
				fmt.Printf("Error reading content.xml: %s", err)
			}

			var content OfficeDoc
			if err := xml.Unmarshal(contentBytes, &content); err != nil {
				fmt.Printf("Error Unmarshal content.xml: %s", err)
			}

			fmt.Println(content)
			var output string
			for _, p := range content.Body.Text.Paragraphs {
				output += p.Content + "\n"
			}

			return output
		}
	}

	return ""
}
