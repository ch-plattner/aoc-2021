package day2

import (
	"errors"
	"github.com/ch-plattner/aoc-2020/util"
	"log"
	"regexp"
	"strconv"
)

const (
	passwordAndPolicyRegex = `(\d+)-(\d+) ([a-zA-Z]): (.*)`
)

type Password string

type Policy struct {
	FirstNum  int64
	SecondNum int64
	Letter    byte
}

type ValidFunc func(Password, Policy) bool

func ProblemOneCalculateValidPasswords(inputFile string) int {
	return calculateValidPasswords(inputFile, isValidPasswordProblemOne)
}

func ProblemTwoCalculateValidPasswords(inputFile string) int {
	return calculateValidPasswords(inputFile, isValidPasswordProblemTwo)
}

func calculateValidPasswords(inputFile string, isValidFunc ValidFunc) int {
	lines, err := util.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var numValid int
	for _, line := range lines {
		policy, password, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		if isValidFunc(password, policy) {
			numValid += 1
		}
	}
	return numValid
}

func isValidPasswordProblemOne(password Password, policy Policy) bool {
	var numAppearances int64
	for i := 0; i < len(password); i++ {
		if password[i] == policy.Letter {
			numAppearances++
		}
	}
	return (numAppearances >= policy.FirstNum) && (numAppearances <= policy.SecondNum)
}

func isValidPasswordProblemTwo(password Password, policy Policy) bool {
	firstPos := policy.FirstNum - 1
	secondPos := policy.SecondNum - 1
	if password[firstPos] == policy.Letter {
		return password[secondPos] != policy.Letter
	}
	return password[secondPos] == policy.Letter
}

func parseLine(line string) (Policy, Password, error) {
	r := regexp.MustCompile(passwordAndPolicyRegex)
	if !r.MatchString(line) {
		return Policy{}, "", errors.New("doesn't match")
	}
	match := r.FindStringSubmatch(line)

	minAppearances, err := strconv.ParseInt(match[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	maxAppearances, err := strconv.ParseInt(match[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	character := match[3]

	return Policy{
		FirstNum:  minAppearances,
		SecondNum: maxAppearances,
		Letter:    character[0],
	}, Password(match[4]), nil
}
