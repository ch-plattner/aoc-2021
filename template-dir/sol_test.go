package day1

import (
	"github.com/ch-plattner/aoc-2020/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	tests := []struct {
		name string
		inputFile string
		expected int
	}{
		{
			name: "basic case",
			inputFile: "mini-input.txt",
			expected: 7,
		},
		{
			name: "input",
			inputFile: util.InputFile,
			expected: 1665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := one(tt.inputFile)
			assert.Equal(t, actual, tt.expected)
		})
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		name string
		inputFile string
		expected int
	}{
		{
			name: "basic case",
			inputFile: "mini-input.txt",
			expected: 7,
		},
		{
			name: "input",
			inputFile: util.InputFile,
			expected: 1665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := two(tt.inputFile)
			assert.Equal(t, actual, tt.expected)
		})
	}
}
