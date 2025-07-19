package txtfiles

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
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

			var output string
			for _, p := range content.Body.Text.Paragraphs {
				output += p.Content + "\n"
			}

			return output
		}
	}

	return ""
}

func WriteOdtFile(outputfile, content string) {
	file, err := os.Create(outputfile)
	if err != nil {
		log.Fatalf("Failed to create output file with name: %s", outputfile)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)

	mimeHeader := &zip.FileHeader{
		Name:   "mimeHeader",
		Method: zip.Store,
	}

	mimeWriter, err := zipWriter.CreateHeader(mimeHeader)
	if err != nil {
		log.Fatalf("Failed to create mimeheader, %s", err)
	}

	_, err = mimeWriter.Write([]byte("application/vnd.oasis.opendocument.text"))
	if err != nil {
		log.Fatalf("Failed to write mimeheader, %s", err)
	}

	contentXML := `<?xml version="1.0" encoding="UTF-8"?>
		<office:document-content 
			xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
			xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0"
			xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0"
			xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0"
			office:version="1.2">
		<office:body>
			<office:text>
			<text:p>` + content + `</text:p>
			</office:text>
		</office:body>
		</office:document-content>`

	writer, err := zipWriter.Create("content.xml")
	if err != nil {
		log.Fatalf("Failed to create content.xml, %s", err)
	}

	_, err = writer.Write([]byte(contentXML))
	if err != nil {
		log.Fatalf("Failed to write content.xml, %s", err)
	}

	stylesXML := `<?xml version="1.0" encoding="UTF-8"?>
		<office:document-styles 
			xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
			xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0"
			xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0"
			office:version="1.2">
		<office:styles/>
		</office:document-styles>`

	style, err := zipWriter.Create("style.xml")
	if err != nil {
		log.Fatalf("Failed to create style.xml, %s", err)
	}

	_, err = style.Write([]byte(stylesXML))
	if err != nil {
		log.Fatalf("Failed to write style.xml, %s", err)
	}

	zipWriter.Close()
}
