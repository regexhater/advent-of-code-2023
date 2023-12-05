package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Number         int
	WinningNumbers []int
	GameNumbers    []int
	Worth          int
}

func (c *Card) CalculateWorth() {
	wCounter := 0
	for _, number := range c.GameNumbers {
		if slices.Contains(c.WinningNumbers, number) {
			wCounter++
		}
	}
	if wCounter > 0 {
		c.Worth = 1
		for i := wCounter - 1; i > 0; i-- {
			c.Worth *= 2
		}
	}
}

func main() {
	fmt.Println("Day 4!")
	cards, err := readInput()
	if err != nil {
		log.Fatalln("could not read input: ", err)
	}
	points := 0
	for _, card := range cards {
		card.CalculateWorth()
		points += card.Worth
	}
	fmt.Printf("The total worth of cards is: %d\n", points)
}

func readInput() ([]Card, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)
	counter := 1
	result := make([]Card, 0)
	for fs.Scan() {
		line := fs.Text()
		lineWithoutHeader := strings.Split(line, ": ")
		if len(lineWithoutHeader) != 2 {
			return nil, fmt.Errorf("could not split line: %s", line)
		}
		winningNormalSplit := strings.Split(lineWithoutHeader[1], " | ")
		winningSplit := strings.Split(
			strings.Replace(
				strings.Trim(winningNormalSplit[0], " "), "  ", " ", -1,
			), " ",
		)
		gameSplit := strings.Split(
			strings.Replace(
				strings.Trim(winningNormalSplit[1], " "), "  ", " ", -1,
			), " ",
		)
		winningNumbers, err := convertNumbers(winningSplit)
		if err != nil {
			return nil, err
		}
		gameNumbers, err := convertNumbers(gameSplit)
		if err != nil {
			return nil, err
		}
		result = append(
			result, Card{
				Number:         counter,
				WinningNumbers: winningNumbers,
				GameNumbers:    gameNumbers,
			},
		)
		counter++
	}

	return result, nil
}

func convertNumbers(numbersAsStrings []string) ([]int, error) {
	result := make([]int, 0)
	for _, ns := range numbersAsStrings {
		number, err := strconv.Atoi(ns)
		if err != nil {
			return nil, err
		}
		result = append(result, number)
	}
	return result, nil
}
