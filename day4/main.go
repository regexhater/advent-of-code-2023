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

const maxCardID = 193

type Card struct {
	Number         int
	WinningNumbers []int
	GameNumbers    []int
	NumberOfHits   int
	Worth          int
}

func (c *Card) CalculateWorth() {
	wCounter := 0
	for _, number := range c.GameNumbers {
		if slices.Contains(c.WinningNumbers, number) {
			wCounter++
		}
	}
	c.NumberOfHits = wCounter
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
	totalCardsWon := 0
	cardsToIds := make(map[int]Card)
	for _, card := range cards {
		card.CalculateWorth()
		points += card.Worth
		cardsToIds[card.Number] = *card
	}
	fmt.Printf("The total worth of cards is: %d\n", points)
	q := make([]Card, 0)
	for _, card := range cards {
		q = append(q, *card)
	}
	for len(q) > 0 {
		card := q[0]
		wonCopies := getWonCopies(card, cardsToIds)
		for _, wonCopy := range wonCopies {
			q = append(q, wonCopy)
		}
		q = q[1:]
		totalCardsWon++
	}
	fmt.Printf("Won %d cards in total\n", totalCardsWon)
}

func readInput() ([]*Card, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)
	counter := 1
	result := make([]*Card, 0)
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
			result, &Card{
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

func getWonCopies(card Card, cardsToIds map[int]Card) []Card {
	copiesWon := make([]Card, 0)
	for i := 1; i <= card.NumberOfHits; i++ {
		wonCardIndex := card.Number + i
		if wonCardIndex > maxCardID {
			break
		}
		copiesWon = append(copiesWon, cardsToIds[wonCardIndex])
	}
	return copiesWon
}
