package main

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

const SUM = 2020

func part1(lines []string) (int, error) {
	numbers, err := utils.ConvertSliceToInts(lines)
	if err != nil {
		return 0, err
	}
	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == SUM {
				return i * j, nil
			}
		}
	}
	return 0, errors.New("no match found")
}

func part2(lines []string) (int, error) {
	numbers, err := utils.ConvertSliceToInts(lines)
	if err != nil {
		return 0, err
	}
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == SUM {
					return i * j * k, nil
				}
			}
		}
	}
	return 0, errors.New("no match found")
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	inputFile := filepath.Join(path.Dir(filename), "input.txt")
	lines := utils.ReadLinesFromFile(inputFile, "\n")

	result, err := part1(lines)
	if err != nil {
		log.Fatalf("Part 1: %s", err)
	}
	fmt.Printf("Result from Part 1: %d\n", result)

	result, err = part2(lines)
	if err != nil {
		log.Fatalf("Part 2: %s", err)
	}
	fmt.Printf("Result from Part 2: %d\n", result)
}
