package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day one!")
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not read file: ", err)
	}
	defer readFile.Close()
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	calibrationSum := 0
	for fs.Scan() {
		lineCalibration, err := findCalibrationValue(fs.Text())
		if err != nil {
			log.Fatalln("Could not read or convert line digits: ", err)
		}
		calibrationSum += lineCalibration
	}
	fmt.Printf("The sum of calibration is: %d\n", calibrationSum)
}

func findCalibrationValue(line string) (int, error) {
	var firstLeftDigit, firstRightDigit rune
	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			firstLeftDigit = ch
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		ch := rune(line[i])
		if ch >= '0' && ch <= '9' {
			firstRightDigit = ch
			break
		}
	}
	sb := strings.Builder{}
	sb.WriteRune(firstLeftDigit)
	sb.WriteRune(firstRightDigit)
	return strconv.Atoi(sb.String())
}

// Read input
// Find first digit from left
// Find first digit from right
// create number and add to sum
