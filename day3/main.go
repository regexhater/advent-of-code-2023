package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X int
	Y int
}

func main() {
	fmt.Println("Day 3!")
	schematics, err := readSchematics()
	if err != nil {
		log.Fatalln("could not read schematics: ", err)
	}
	sumOfPartNumbers := 0
	gearsCoordsToAdjacentNumbers := make(map[Coords]*[]int)
	for i := 0; i < len(schematics); i++ {
		foundNumber := false
		numberStart := 0
		numberEnd := 0
		for j := 0; j < len(schematics[i]); j++ {
			ch := schematics[i][j]
			if ch >= '0' && ch <= '9' {
				if !foundNumber {
					foundNumber = true
					numberStart = j
				} else if foundNumber && j == len(schematics[i])-1 {
					numberEnd = j
					foundNumber = false
					if isAdjacentToSymbol(schematics, i, numberStart, numberEnd) {
						num := convertToNumber(schematics, i, numberStart, numberEnd)
						sumOfPartNumbers += num
						cords := getCoordsIfAdjacentToGear(schematics, i, numberStart, numberEnd)
						if cords != nil {
							adjacentNums, ok := gearsCoordsToAdjacentNumbers[*cords]
							if !ok {
								gearsCoordsToAdjacentNumbers[*cords] = &[]int{num}
							} else {
								*adjacentNums = append(*adjacentNums, num)
							}
						}
					}
				}
			} else {
				if foundNumber {
					numberEnd = j - 1
					if isAdjacentToSymbol(schematics, i, numberStart, numberEnd) {
						num := convertToNumber(schematics, i, numberStart, numberEnd)
						sumOfPartNumbers += num
						cords := getCoordsIfAdjacentToGear(schematics, i, numberStart, numberEnd)
						if cords != nil {
							adjacentNums, ok := gearsCoordsToAdjacentNumbers[*cords]
							if !ok {
								gearsCoordsToAdjacentNumbers[*cords] = &[]int{num}
							} else {
								*adjacentNums = append(*adjacentNums, num)
							}
						}
					}
				}
				foundNumber = false
			}
		}
	}
	fmt.Printf("The sum of all part numbers is: %d\n", sumOfPartNumbers)
	sumOfGears := 0
	for _, nums := range gearsCoordsToAdjacentNumbers {
		if len(*nums) == 2 {
			sumOfGears += (*nums)[0] * (*nums)[1]
		}
	}
	fmt.Printf("The sum of all gear ratios if %d\n", sumOfGears)
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
		if (i-1 >= 0 && isSymbol(schematics[x][i-1])) && i-1 < startY ||
			(i-1 >= 0 && x-1 >= 0 && isSymbol(schematics[x-1][i-1])) ||
			(x-1 >= 0 && isSymbol(schematics[x-1][i])) ||
			(x-1 >= 0 && i+1 < len(schematics[x]) && isSymbol(schematics[x-1][i+1])) ||
			(i+1 < len(schematics[x]) && isSymbol(schematics[x][i+1])) && i+1 > endY ||
			(x+1 < len(schematics) && i+1 < len(schematics[x]) && isSymbol(schematics[x+1][i+1])) ||
			(x+1 < len(schematics) && isSymbol(schematics[x+1][i])) ||
			(x+1 < len(schematics) && i-1 >= 0 && isSymbol(schematics[x+1][i-1])) {
			return true
		}
	}
	return false
}

func getCoordsIfAdjacentToGear(schematics []string, x int, startY int, endY int) *Coords {
	for i := startY; i <= endY; i++ {
		if i-1 >= 0 && schematics[x][i-1] == '*' {
			return &Coords{
				X: x,
				Y: i - 1,
			}
		} else if i-1 >= 0 && x-1 >= 0 && schematics[x-1][i-1] == '*' {
			return &Coords{
				X: x - 1,
				Y: i - 1,
			}
		} else if x-1 >= 0 && schematics[x-1][i] == '*' {
			return &Coords{
				X: x - 1,
				Y: i,
			}
		} else if x-1 >= 0 && i+1 < len(schematics[x]) && schematics[x-1][i+1] == '*' {
			return &Coords{
				X: x - 1,
				Y: i + 1,
			}
		} else if i+1 < len(schematics[x]) && schematics[x][i+1] == '*' {
			return &Coords{
				X: x,
				Y: i + 1,
			}
		} else if x+1 < len(schematics) && i+1 < len(schematics[x]) && schematics[x+1][i+1] == '*' {
			return &Coords{
				X: x + 1,
				Y: i + 1,
			}
		} else if x+1 < len(schematics) && schematics[x+1][i] == '*' {
			return &Coords{
				X: x + 1,
				Y: i,
			}
		} else if x+1 < len(schematics) && i-1 >= 0 && schematics[x+1][i-1] == '*' {
			return &Coords{
				X: x + 1,
				Y: i - 1,
			}
		}
	}
	return nil
}

func isSymbol(ch uint8) bool {
	return ch != '.'
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
