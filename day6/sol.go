package day1

import (
	"github.com/ch-plattner/aoc-2021/util"
	"strconv"
	"strings"
)

const (
	newbornAge     = 8
	afterBirthAge  = 6
	givingBirthAge = 0
)

func one(inputFile string, days int) (int, error) {
	fishes, err := extractFish(inputFile)
	if err != nil {
		return 0, err
	}
	for i := 0; i < days; i++ {
		numFishInDay := len(fishes)
		for f := 0; f < numFishInDay; f++ {
			newAge, maybeBaby := processFish(fishes[f])
			fishes[f] = newAge
			if maybeBaby != nil {
				fishes = append(fishes, *maybeBaby)
			}
		}
	}
	return len(fishes), nil
}

func extractFish(inputFile string) ([]int, error) {
	rawFile, err := util.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	rawFish := strings.Split(rawFile[0], ",")
	var fishes []int
	for _, s := range rawFish {
		fish, err := strconv.ParseInt(s,10, 64)
		if err != nil {
			return nil, err
		}
		fishes = append(fishes, int(fish))
	}
	return fishes, nil
}

func processFish(dayTimer int) (int, *int) {
	if dayTimer == 0 {
		babyFish := newbornAge
		return afterBirthAge, &babyFish
	}
	return dayTimer-1, nil
}

func two(inputFile string, days int) (int, error) {
	fishes, err := extractFish(inputFile)
	if err != nil {
		return 0, err
	}
	// all lanterns produced on the same day will have the same age
	fishByAge := make(map[int]uint, 0)
	for _, age := range fishes {
		fishByAge[age] += 1
	}
	for i := 0; i < days; i++ {
		fishGivingBirth := fishByAge[givingBirthAge]
		newFishByAge := map[int]uint {
			newbornAge: fishGivingBirth,
		}
		for age := 1; age <= newbornAge; age++ {
			newFishByAge[age-1] = fishByAge[age]
		}
		newFishByAge[afterBirthAge] += fishGivingBirth
		fishByAge = newFishByAge
	}
	totalFish := 0
	for age := range fishByAge {
		totalFish += int(fishByAge[age])
	}
	return totalFish, nil
}
