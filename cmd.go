package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type CmdFlags struct {
	ConvertImg string
	ConvertTxt string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.ConvertImg, "img", "", "Convert a format of a image. Select the formats jpeg, jpg and png")
	flag.StringVar(&cf.ConvertTxt, "txt", "", "Convert a format of a txt")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute() {
	switch {
	case cf.ConvertImg != "":
		fileFrom, extensionFrom, fileTo, extensionTo := cf.separateStringFiles(cf.ConvertImg)

		ConvertImg(fileFrom, extensionFrom, fileTo, extensionTo)
	case cf.ConvertTxt != "":
		fileFrom, extensionFrom, fileTo, extensionTo := cf.separateStringFiles(cf.ConvertTxt)

		ConvertTxt(fileFrom, extensionFrom, fileTo, extensionTo)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

func (cf *CmdFlags) separateStringFiles(parts string) (string, string, string, string) {
	s := strings.Split(parts, ":")

	if len(s) != 2 {
		fmt.Println("Error, invalid format for add. Please use inputfile:outputfile")
		os.Exit(1)
	}
	fileFrom := strings.Split(s[0], ".")
	fileTo := strings.Split(s[1], ".")

	return s[0], fileFrom[1], s[1], fileTo[1]
}
