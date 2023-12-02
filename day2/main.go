package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id       int
	MaxRed   int
	MaxGreen int
	MaxBlue  int
	Rounds   []Round
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

func (g *Game) findMaximum() {
	for _, round := range g.Rounds {
		g.MaxRed = max(g.MaxRed, round.Red)
		g.MaxGreen = max(g.MaxGreen, round.Green)
		g.MaxBlue = max(g.MaxBlue, round.Blue)
	}
}

func (g *Game) isPossible(red int, green int, blue int) bool {
	return g.MaxRed <= red &&
		g.MaxGreen <= green &&
		g.MaxBlue <= blue
}

func main() {
	fmt.Println("Day two!")
	games, err := readGames()
	if err != nil {
		log.Fatalln("could not read games: ", err)
	}
	idSum := 0
	for _, game := range games {
		game.findMaximum()
		if game.isPossible(12, 13, 14) {
			idSum += game.Id
		}
	}
	fmt.Printf("Sum of id equals: %d\n", idSum)
}

func readGames() ([]Game, error) {
	result := make([]Game, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	counter := 1
	for fs.Scan() {
		line := fs.Text()
		// split game id from rounds
		gameSplit := strings.Split(line, ": ")
		if len(gameSplit) != 2 {
			return nil, errors.New("could not split line: " + line)
		}
		// split rounds
		roundSplit := strings.Split(gameSplit[1], "; ")
		rounds := make([]Round, 0)
		for _, roundString := range roundSplit {
			round := Round{}
			// split pulls within round
			pullsSplit := strings.Split(roundString, ", ")
			for _, pull := range pullsSplit {
				// split number and color
				wordSplit := strings.Split(pull, " ")
				if len(wordSplit) != 2 {
					return nil, errors.New("could not split words: " + pull)
				}
				num, err := strconv.Atoi(wordSplit[0])
				if err != nil {
					return nil, err
				}
				switch wordSplit[1] {
				case "red":
					round.Red += num
				case "green":
					round.Green += num
				case "blue":
					round.Blue += num
				}
			}
			rounds = append(rounds, round)
		}
		result = append(
			result, Game{
				Id:     counter,
				Rounds: rounds,
			},
		)
		counter++
	}

	return result, nil
}
