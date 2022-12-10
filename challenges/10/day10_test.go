package main

import (
	"testing"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePart01(t *testing.T) {
	t.Parallel()
	input := utils.LoadPuzzleInput("example.txt")
	expected := 13140

	actual := SolvePart01(input)
	assert.Equal(t, expected, actual)
}

func Test_GetXatCycle(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "cycle 20", input: 20, expected: 420},
		{name: "cycle 60", input: 60, expected: 1140},
		{name: "cycle 100", input: 100, expected: 1800},
		{name: "cycle 140", input: 140, expected: 2940},
		{name: "cycle 180", input: 180, expected: 2880},
		{name: "cycle 220", input: 220, expected: 3960},
	}

	input := utils.LoadPuzzleInput("example.txt")
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := getXatCycle(input, test.input) * test.input
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_SolvePart02(t *testing.T) {
	t.Parallel()
	input := utils.LoadPuzzleInput("example.txt")
	expected := "##..##..##..##..##..##..##..##..##..##.." +
		"###...###...###...###...###...###...###." +
		"####....####....####....####....####...." +
		"#####.....#####.....#####.....#####....." +
		"######......######......######......####" +
		"#######.......#######.......#######....."

	actual := SolvePart02(input)
	assert.Equal(t, expected, actual)
}
