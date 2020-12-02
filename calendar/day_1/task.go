package main

import (
	"errors"
	"fmt"
	"github.com/Jakob-em/advent-of-code-2020/utils"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

const SUM = 2020

func Part1(lines []string) (int, error) {
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

func Part2(lines []string) (int, error) {
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

	result, err := Part1(lines)
	if err != nil {
		log.Fatalf("Part 1: %s", err)
	}
	fmt.Printf("Result from Part 1: %d\n", result)

	result, err = Part2(lines)
	if err != nil {
		log.Fatalf("Part 2: %s", err)
	}
	fmt.Printf("Result from Part 2: %d\n", result)
}
