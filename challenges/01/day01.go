package main

import (
	"log"
	"sort"
	"strconv"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
)

func main() {
	input := utils.LoadPuzzleInput("challenges/01/input.txt")
	log.Printf("Soluce part 1 is %d", SolvePart01(input))
	log.Printf("Soluce part 2 is %d", SolvePart02(input))
}

func SolvePart01(input []string) int {
	elves := compute(input)
	return getTotalByElf(elves)[0]
}

func SolvePart02(input []string) int {
	elves := compute(input)
	total := getTotalByElf(elves)
	top3Total := 0
	for i := 0; i < 3; i++ {
		top3Total += total[i]
	}
	return top3Total
}

func compute(input []string) [][]int {
	var elves [][]int
	var elf []int
	for _, line := range input {
		if line == "" {
			elves = append(elves, elf)
			elf = []int{}
			continue
		}
		food, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		elf = append(elf, food)
	}
	return elves
}

func getTotalByElf(elves [][]int) []int {
	var total []int
	for _, food := range elves {
		elfTotalFood := getElfTotalFood(food)
		total = append(total, elfTotalFood)
	}
	sort.Slice(total, func(i, j int) bool {
		return total[i] > total[j]
	})
	return total
}

func getElfTotalFood(elf []int) int {
	total := 0
	for _, food := range elf {
		total += food
	}
	return total
}
