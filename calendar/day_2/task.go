package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type lineData struct {
	password  string
	character string
	min       int
	max       int
}

func part1(lines []string) (int, error) {
	valid := 0
	for _, line := range lines {
		lineData := parseLine(line)
		count := strings.Count(lineData.password, lineData.character)
		if count >= lineData.min && count <= lineData.max {
			valid = valid + 1
		}
	}

	return valid, nil
}

func part2(lines []string) (int, error) {
	valid := 0
	for _, line := range lines {
		lineData := parseLine(line)

		isMinMatch := lineData.password[lineData.min-1] == lineData.character[0]
		isMaxMatch := lineData.password[lineData.max-1] == lineData.character[0]

		if (isMinMatch || isMaxMatch) && !(isMinMatch && isMaxMatch) {
			valid = valid + 1
		}
	}

	return valid, nil
}

func parseLine(line string) lineData {
	parts := strings.Split(line, ":")
	password := strings.Trim(parts[1], " ")

	rangeAndChar := strings.Split(parts[0], " ")

	allowedRange := strings.Split(rangeAndChar[0], "-")
	min, _ := strconv.Atoi(allowedRange[0])
	max, _ := strconv.Atoi(allowedRange[1])

	return lineData{password: password, character: rangeAndChar[1], min: min, max: max}

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
