package util

import (
	"bufio"
	"log"
	"os"
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
