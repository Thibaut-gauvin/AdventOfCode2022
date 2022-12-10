package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LoadPuzzleInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	_ = file.Close()

	return lines
}

func StrToInt(val string) int {
	value, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
