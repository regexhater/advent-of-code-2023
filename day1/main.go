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
	sb := strings.Builder{}
	for _, ch := range line {
		dig := checkForDigit(ch, sb.String(), true)
		if dig != nil {
			firstLeftDigit = *dig
			sb.Reset()
			break
		}
		sb.WriteRune(ch)
	}
	for i := len(line) - 1; i >= 0; i-- {
		ch := rune(line[i])
		dig := checkForDigit(ch, sb.String(), false)
		if dig != nil {
			firstRightDigit = *dig
			sb.Reset()
			break
		}
		temp := sb.String()
		sb.Reset()
		sb.WriteRune(ch)
		sb.WriteString(temp)
	}
	sb.WriteRune(firstLeftDigit)
	sb.WriteRune(firstRightDigit)
	s := sb.String()
	return strconv.Atoi(s)
}

func checkForDigit(ch rune, lineSoFar string, isFromLeft bool) *rune {
	if ch >= '0' && ch <= '9' {
		return &ch
	}
	sb := strings.Builder{}
	if isFromLeft {
		sb.WriteString(lineSoFar)
		sb.WriteRune(ch)
	} else {
		sb.WriteRune(ch)
		sb.WriteString(lineSoFar)
	}
	var dig rune
	potentialSpelledDigit := sb.String()
	if strings.Contains(potentialSpelledDigit, "one") {
		dig = '1'
	} else if strings.Contains(potentialSpelledDigit, "two") {
		dig = '2'
	} else if strings.Contains(potentialSpelledDigit, "three") {
		dig = '3'
	} else if strings.Contains(potentialSpelledDigit, "four") {
		dig = '4'
	} else if strings.Contains(potentialSpelledDigit, "five") {
		dig = '5'
	} else if strings.Contains(potentialSpelledDigit, "six") {
		dig = '6'
	} else if strings.Contains(potentialSpelledDigit, "seven") {
		dig = '7'
	} else if strings.Contains(potentialSpelledDigit, "eight") {
		dig = '8'
	} else if strings.Contains(potentialSpelledDigit, "nine") {
		dig = '9'
	} else {
		return nil
	}
	return &dig
}
