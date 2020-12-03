package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type movement struct {
	x int
	y int
}

func part1(lines []string) (int, error) {
	encounteredTrees := countEncounteredTrees(lines, movement{
		x: 3,
		y: 1,
	})

	return encounteredTrees, nil
}

func countEncounteredTrees(lines []string, move movement) int {
	encounteredTrees := 0
	posX := 0
	for i := 0; i < len(lines); i = i + move.y {
		line := lines[i]
		if isTreeAtPosition(line, posX) {
			encounteredTrees++
		}
		posX += move.x
	}
	return encounteredTrees
}

func isTreeAtPosition(line string, posX int) bool {
	return line[posX%len(line)] == '#'
}

var movesToMultiply = []movement{
	{x: 1, y: 1},
	{x: 3, y: 1},
	{x: 5, y: 1},
	{x: 7, y: 1},
	{x: 1, y: 2},
}

func part2(lines []string) (int, error) {
	result := 1
	for _, move := range movesToMultiply {
		result *= countEncounteredTrees(lines, move)
	}
	return result, nil
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
