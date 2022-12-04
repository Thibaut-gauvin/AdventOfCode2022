package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
	"golang.org/x/exp/slices"
)

func main() {
	input := utils.LoadPuzzleInput("challenges/04/input.txt")
	log.Printf("Soluce part 1 is %d", SolvePart01(input))
	log.Printf("Soluce part 2 is %d", SolvePart02(input))
}

func SolvePart01(input []string) int {
	fullContain := 0
	for _, line := range input {
		pairs := strings.Split(line, ",")
		elf1, elf2 := pairs[0], pairs[1]
		elf1Sections := strings.Split(elf1, "-")
		elf2Sections := strings.Split(elf2, "-")
		if isFullyContain(
			strToInt(elf1Sections[0]), strToInt(elf1Sections[1]),
			strToInt(elf2Sections[0]), strToInt(elf2Sections[1]),
		) {
			fullContain++
		}
	}
	return fullContain
}

func SolvePart02(input []string) int {
	overlap := 0
	for _, line := range input {
		pairs := strings.Split(line, ",")
		elf1, elf2 := pairs[0], pairs[1]
		elf1Sections := strings.Split(elf1, "-")
		elf2Sections := strings.Split(elf2, "-")
		if isOverlapping(elf1Sections, elf2Sections) {
			overlap++
		}
	}
	return overlap
}

func strToInt(item string) int {
	value, err := strconv.Atoi(item)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func isFullyContain(elf1section1, elf1section2, elf2section1, elf2section2 int) bool {
	fullyContain := false
	if (elf1section1 >= elf2section1 && elf1section2 <= elf2section2) ||
		(elf1section1 <= elf2section1 && elf1section2 >= elf2section2) {
		fullyContain = true
	}
	return fullyContain
}

func isOverlapping(elf1Sections, elf2Sections []string) bool {
	var interval1 []int
	for i := strToInt(elf1Sections[0]); i <= strToInt(elf1Sections[1]); i++ {
		interval1 = append(interval1, i)
	}

	var interval2 []int
	for i := strToInt(elf2Sections[0]); i <= strToInt(elf2Sections[1]); i++ {
		interval2 = append(interval2, i)
	}

	overlaps := false
	for _, section := range interval2 {
		if slices.Contains(interval1, section) {
			overlaps = true
		}
	}
	return overlaps
}
