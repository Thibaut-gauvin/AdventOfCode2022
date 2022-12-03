package main

import (
	"log"
	"strings"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
)

func main() {
	input := utils.LoadPuzzleInput("challenges/03/input.txt")
	log.Printf("Soluce part 1 is %d", SolvePart01(input))
	log.Printf("Soluce part 2 is %d", SolvePart02(input))
}

func SolvePart01(input []string) int {
	var duplicates string
	for _, line := range input {
		compartment1, compartment2 := line[0:len(line)/2], line[len(line)/2:]
		duplicates += findDuplicatesInCompartments(compartment1, compartment2)
	}
	return getPriorityScore(duplicates)
}

func SolvePart02(input []string) int {
	var duplicates string
	for i := 0; i < len(input); i += 3 {
		duplicates += findDuplicatesInGroup(input[i : i+3])
	}
	return getPriorityScore(duplicates)
}

func findDuplicatesInCompartments(compartment1, compartment2 string) string {
	var duplicates string
	for _, item1 := range compartment1 {
		item1Str := string(item1)
		isPresent := strings.Contains(compartment2, item1Str)
		if isPresent {
			if !strings.Contains(duplicates, item1Str) {
				duplicates += item1Str
			}
		}
	}
	return duplicates
}

func findDuplicatesInGroup(elvesRucksacks []string) string {
	itemList := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var duplicates string
	for _, char := range itemList {
		charStr := string(char)
		if strings.Contains(elvesRucksacks[0], charStr) &&
			strings.Contains(elvesRucksacks[1], charStr) &&
			strings.Contains(elvesRucksacks[2], charStr) {
			duplicates += charStr
		}
	}
	return duplicates
}

func getPriorityScore(duplicateItem string) int {
	scoreMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13,
		"n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25,
		"z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37,
		"L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49,
		"X": 50, "Y": 51, "Z": 52,
	}
	var score int
	for _, item := range duplicateItem {
		score += scoreMap[string(item)]
	}
	return score
}
