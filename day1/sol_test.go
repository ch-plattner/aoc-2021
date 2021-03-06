package day1

import (
	"github.com/ch-plattner/aoc-2020/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	tests := []struct {
		name string
		nums []int64
		expected int
	}{
		{
			name: "basic case",
			nums: []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expected: 7,
		},
		{
			name: "input",
			expected: 1665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			if len(nums) == 0 {
				input, err := util.ReadIntFile(util.InputFile)
				nums = append(nums, input...)
				assert.NoError(t, err)
			}
			actual := one(nums)
			assert.Equal(t, actual, tt.expected)

		})
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		name string
		nums []int64
		expected int
	}{
		{
			name: "basic case",
			nums: []int64{607,618,618,617,647,716,769,792},
			expected: 5,
		},
		{
			name: "input",
			expected: 1702,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			if len(nums) == 0 {
				input, err := util.ReadIntFile(util.InputFile)
				nums = append(nums, input...)
				assert.NoError(t, err)
			}
			actual := two(nums)
			assert.Equal(t, actual, tt.expected)
		})
	}
}
