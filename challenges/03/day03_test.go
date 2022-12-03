package main

import (
	"testing"

	"github.com/Thibaut-gauvin/AdventOfCode2022/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SolvePart01(t *testing.T) {
	t.Parallel()
	input := utils.LoadPuzzleInput("example.txt")
	expected := 157

	result := SolvePart01(input)
	assert.Equal(t, expected, result)
}

func Test_SolvePart02(t *testing.T) {
	t.Parallel()
	input := utils.LoadPuzzleInput("example.txt")
	expected := 70

	result := SolvePart02(input)
	assert.Equal(t, expected, result)
}
