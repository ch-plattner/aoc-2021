package day1

import (
	"github.com/ch-plattner/aoc-2021/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	tests := []struct {
		name string
		inputFile string
		days int
		expected int
	}{
		{
			name: "basic case",
			inputFile: "mini-input.txt",
			days: 18,
			expected: 26,
		},
		{
			name: "basic case",
			inputFile: "mini-input.txt",
			days: 80,
			expected: 5934,
		},
		{
			name: "input",
			inputFile: util.InputFile,
			days: 80,
			expected: 350917,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := one(tt.inputFile, tt.days)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		name string
		inputFile string
		days int
		expected int
	}{
		{
			name: "basic case - 18 days",
			inputFile: "mini-input.txt",
			days: 18,
			expected: 26,
		},
		{
			name: "basic case - 80 days",
			inputFile: "mini-input.txt",
			days: 80,
			expected: 5934,
		},
		{
			name: "basic case - 256 days",
			inputFile: "mini-input.txt",
			days: 256,
			expected: 26984457539,
		},
		{
			name: "input",
			inputFile: util.InputFile,
			days: 256,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := two(tt.inputFile, tt.days)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
