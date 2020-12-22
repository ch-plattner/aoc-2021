package day1

import (
	"fmt"
	"testing"
)

func Test_ProblemOneFind2020Sum(t *testing.T) {
	sol := ProblemOneFind2020Sum("./input1.txt")
	fmt.Println(sol)
}

func Test_ProblemTwoFind2020Sum(t *testing.T) {
	sol := ProblemTwoFind2020Sum("./input1.txt")
	fmt.Println(sol)
}