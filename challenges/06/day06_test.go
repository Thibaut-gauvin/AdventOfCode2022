package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolvePart01(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "valid 1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 7,
		},
		{
			name:     "valid 2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			name:     "valid 3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			name:     "valid 4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			name:     "valid 5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := SolvePart01(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_SolvePart02(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "valid 1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			name:     "valid 2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			name:     "valid 3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			name:     "valid 4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			name:     "valid 5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := SolvePart02(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
