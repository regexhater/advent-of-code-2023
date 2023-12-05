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
	fmt.Println("Day 3!")
	schematics, err := readSchematics()
	if err != nil {
		log.Fatalln("could not read schematics: ", err)
	}
	fmt.Printf("Shematics dimensions: %d x %d\n", len(schematics), len(schematics[0]))
	sumOfPartNumbers := 0
	takenNumbers := make([]int, 0)
	skippedNumber := make([]int, 0)
	for i := 0; i < len(schematics); i++ {
		foundDigit := false
		digitStart := 0
		digitEnd := 0
		for j := 0; j < len(schematics[i]); j++ {
			ch := schematics[i][j]
			if ch >= '0' && ch <= '9' {
				// fmt.Println("It's a digit")
				if !foundDigit {
					foundDigit = true
					digitStart = j
				}
			} else {
				if foundDigit {
					digitEnd = j - 1
					if isAdjacentToSymbol(schematics, i, digitStart, digitEnd) {
						num := convertToNumber(schematics, i, digitStart, digitEnd)
						sumOfPartNumbers += num
						takenNumbers = append(takenNumbers, num)
					} else {
						num := convertToNumber(schematics, i, digitStart, digitEnd)
						skippedNumber = append(skippedNumber, num)
					}
				}
				foundDigit = false
			}
		}
	}
	fmt.Println(skippedNumber)
	fmt.Printf("The sum of all part numbers is: %d\n", sumOfPartNumbers)
}

func readSchematics() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)
	schematics := make([]string, 0)
	for fs.Scan() {
		schematics = append(schematics, fs.Text())
	}
	return schematics, nil
}

func isAdjacentToSymbol(schematics []string, x int, startY int, endY int) bool {
	for i := startY; i <= endY; i++ {
		if (i-1 >= 0 && isSymbol(schematics[x][i-1])) ||
			(i-1 >= 0 && x-1 >= 0 && isSymbol(schematics[x-1][i-1])) ||
			(x-1 >= 0 && isSymbol(schematics[x-1][i])) ||
			(x-1 >= 0 && i+1 < len(schematics[x]) && isSymbol(schematics[x-1][i+1])) ||
			(i+1 < len(schematics[x]) && isSymbol(schematics[x][i+1])) ||
			(x+1 < len(schematics) && i+1 < len(schematics[x]) && isSymbol(schematics[x+1][i+1])) ||
			(x+1 < len(schematics) && isSymbol(schematics[x+1][i])) ||
			(x+1 < len(schematics) && i-1 >= 0 && isSymbol(schematics[x+1][i-1])) {
			return true
		}
	}
	return false
}

func isSymbol(ch uint8) bool {
	return ch != '.' && (ch < '0' || ch > '9')
}

func convertToNumber(schematics []string, x int, start int, end int) int {
	sb := strings.Builder{}
	for i := start; i <= end; i++ {
		sb.WriteRune(rune(schematics[x][i]))
	}
	str := sb.String()
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln("could not convert to number: ", start)
	}
	return number
}
