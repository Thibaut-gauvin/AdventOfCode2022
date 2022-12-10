package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
)

var getXExpr = regexp.MustCompile(`^addx (-?[0-9]*)$`)

func main() {
	input := utils.LoadPuzzleInput("challenges/10/input.txt")
	log.Printf("Soluce part 1 is %d", SolvePart01(input))

	crt := SolvePart02(input)
	log.Print("Soluce part 2 is\n")
	displayCRT(crt)
}

func SolvePart01(inputs []string) int {
	signal := 0
	signal += getXatCycle(inputs, 20) * 20
	signal += getXatCycle(inputs, 60) * 60
	signal += getXatCycle(inputs, 100) * 100
	signal += getXatCycle(inputs, 140) * 140
	signal += getXatCycle(inputs, 180) * 180
	signal += getXatCycle(inputs, 220) * 220

	return signal
}

func SolvePart02(inputs []string) string {
	crt := ""
	for cycle := 1; cycle <= 240; cycle++ {
		cpuX := getXatCycle(inputs, cycle)
		sprite := generateSprite(cpuX)
		var crtPos int
		switch {
		case cycle <= 40:
			crtPos = cycle - 1
		case cycle > 40 && cycle <= 80:
			crtPos = cycle - 40 - 1
		case cycle > 80 && cycle <= 120:
			crtPos = cycle - 80 - 1
		case cycle > 120 && cycle <= 160:
			crtPos = cycle - 120 - 1
		case cycle > 160 && cycle <= 200:
			crtPos = cycle - 160 - 1
		case cycle > 200 && cycle <= 240:
			crtPos = cycle - 200 - 1
		}
		if sprite[crtPos] == '#' {
			crt += "#"
		} else {
			crt += "."
		}
	}
	return crt
}

func getXatCycle(inputs []string, cycle int) int {
	cpuX := 1
	currentCycle := 0

	for _, input := range inputs {
		if strings.Contains(input, "noop") {
			currentCycle++
		}
		if currentCycle == cycle {
			return cpuX
		}
		if strings.Contains(input, "addx") {
			currentCycle++
			if currentCycle == cycle {
				return cpuX
			}
			currentCycle++
			if currentCycle == cycle {
				return cpuX
			}
			match := getXExpr.FindStringSubmatch(input)
			cpuX += utils.StrToInt(match[1])
		}
	}
	return cpuX
}

func generateSprite(pos int) []rune {
	sprite := []rune("........................................")
	switch {
	case pos <= 0:
		sprite[0] = '#'
		sprite[1] = '#'
	case pos > 0 && pos < len(sprite)-1:
		sprite[pos-1] = '#'
		sprite[pos] = '#'
		sprite[pos+1] = '#'
	}
	return sprite
}

//nolint:forbidigo
func displayCRT(crt string) {
	if len(crt) != 240 {
		log.Fatalf("Length of crt is out of bound :'(\n (len %d)", len(crt))
	}

	fmt.Println(crt[0:40])
	fmt.Println(crt[40:80])
	fmt.Println(crt[80:120])
	fmt.Println(crt[120:160])
	fmt.Println(crt[160:200])
	fmt.Println(crt[200:240])
}
