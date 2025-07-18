package txtfiles

import (
	"bufio"
	"log"
	"os"
)

func ReadTxtFile(inputfile string) string {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line += scanner.Text() + "\n"
	}
	return line
}
