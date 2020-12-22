package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ProblemOneFind2020Sum(inputFile string) int64 {
	nums, err := readFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == 2020 {
				return nums[i]*nums[j]
			}
		}
	}
	return 0
}

func ProblemTwoFind2020Sum(inputFile string) int64 {
	nums, err := readFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 2020 {
					return nums[i]*nums[j]*nums[k]
				}
			}
		}
	}
	return 0
}

func readFile(inputFile string) ([]int64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var res []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawLine := scanner.Text()
		num, err := strconv.ParseInt(rawLine, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res, file.Close()
}
