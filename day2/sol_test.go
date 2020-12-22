package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProblemOneCalculateValidPasswords(t *testing.T) {
	sol := ProblemOneCalculateValidPasswords("./sample.txt")
	assert.Equal(t, 2, sol)

	sol = ProblemOneCalculateValidPasswords("./input.txt")
	fmt.Println(sol)
}

func TestProblemTwoCalculateValidPasswords(t *testing.T) {
	sol := ProblemTwoCalculateValidPasswords("./sample.txt")
	assert.Equal(t, 1, sol)

	sol = ProblemTwoCalculateValidPasswords("./input.txt")
	fmt.Println(sol)
}