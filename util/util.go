package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var res []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res, file.Close()
}

func ReadIntFile(inputFile string) ([]int64, error) {
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