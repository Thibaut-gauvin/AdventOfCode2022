package main

import (
	"log"
	"strings"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
)

func main() {
	input := utils.LoadPuzzleInput("challenges/02/input.txt")
	log.Printf("Soluce part 1 is %d", SolvePart01(input))
	log.Printf("Soluce part 2 is %d", SolvePart02(input))
}

func SolvePart01(input []string) int {
	score := 0
	for _, line := range input {
		match := strings.Split(line, " ")
		score += getScore(match)
	}

	return score
}

func SolvePart02(input []string) int {
	score := 0
	for _, line := range input {
		match := strings.Split(line, " ")
		score += getScoreWithIntention(match)
	}

	return score
}

func getScore(match []string) int {
	var score int
	player1, player2 := match[0], match[1]

	if player1 == "A" && player2 == "X" {
		score = 3 + 1
	}
	if player1 == "A" && player2 == "Y" {
		score = 6 + 2
	}
	if player1 == "A" && player2 == "Z" {
		score = 0 + 3
	}

	if player1 == "B" && player2 == "X" {
		score = 0 + 1
	}
	if player1 == "B" && player2 == "Y" {
		score = 3 + 2
	}
	if player1 == "B" && player2 == "Z" {
		score = 6 + 3
	}

	if player1 == "C" && player2 == "X" {
		score = 6 + 1
	}
	if player1 == "C" && player2 == "Y" {
		score = 0 + 2
	}
	if player1 == "C" && player2 == "Z" {
		score = 3 + 3
	}

	return score
}

func getScoreWithIntention(match []string) int {
	var score int
	player1, player2 := match[0], match[1]

	if player1 == "A" && player2 == "X" {
		score = 0 + 3
	}
	if player1 == "A" && player2 == "Y" {
		score = 3 + 1
	}
	if player1 == "A" && player2 == "Z" {
		score = 6 + 2
	}

	if player1 == "B" && player2 == "X" {
		score = 0 + 1
	}
	if player1 == "B" && player2 == "Y" {
		score = 3 + 2
	}
	if player1 == "B" && player2 == "Z" {
		score = 6 + 3
	}

	if player1 == "C" && player2 == "X" {
		score = 0 + 2
	}
	if player1 == "C" && player2 == "Y" {
		score = 3 + 3
	}
	if player1 == "C" && player2 == "Z" {
		score = 6 + 1
	}

	return score
}
