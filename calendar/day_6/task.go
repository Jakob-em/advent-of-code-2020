package main

import (
	"fmt"
	"github.com/Jakob-em/advent-of-code-2020/utils"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func part1(lines []string) (int, error) {

	total := 0
	for _, line := range lines {
		cleaned := strings.Replace(line, "\n", "", -1)
		questions := map[rune]bool{}
		for _, c := range cleaned {
			questions[c] = true
		}
		total += len(questions)
	}

	return total, nil
}

func part2(lines []string) (int, error) {
	total := 0
	for _, line := range lines {
		persons := strings.Count(line, "\n") + 1
		questions := map[rune]int{}

		for _, c := range line {
			questions[c] += 1
			if questions[c] == persons {
				total += 1
			}
		}
	}

	return total, nil
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	inputFile := filepath.Join(path.Dir(filename), "input.txt")
	lines := utils.ReadLinesFromFile(inputFile, "\n\n")

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
