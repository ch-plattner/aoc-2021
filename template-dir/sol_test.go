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
			expected: 0,
		},
		{
			name: "input",
			inputFile: util.InputFile,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := one(tt.inputFile)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
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
			expected: 0,
		},
		{
			name: "input",
			inputFile: util.InputFile,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := two(tt.inputFile)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
