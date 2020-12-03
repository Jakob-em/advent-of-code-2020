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

func part1(lines []string) (int, error) {
	return 0, errors.New("not implemented")
}

func part2(lines []string) (int, error) {
	return 0, errors.New("not implemented")
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
