package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
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

func ConvertTxt(inputfile, outputfile string) {
	fmt.Println(inputfile)
	fmt.Println(outputfile)
}
