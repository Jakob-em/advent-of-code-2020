package main

import (
	"errors"
	"fmt"
	"github.com/Jakob-em/advent-of-code-2020/utils"
	"log"
	"math"
	"path"
	"path/filepath"
	"runtime"
)

func part1(lines []string) (int, error) {
	highestId := 0
	for _, line := range lines {
		currentId := findSeatId(line)
		if currentId > highestId {
			highestId = currentId
		}
	}

	return highestId, nil
}

func part2(lines []string) (int, error) {
	seats := map[int]bool{}
	for _, line := range lines {
		seats[findSeatId(line)] = true
	}

	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			seatId := id(row, col)
			_, left := seats[seatId-1]
			_, right := seats[seatId+1]
			_, occupied := seats[seatId]
			if left && right && !occupied {
				return seatId, nil
			}
		}
	}
	return -1, errors.New("could not find empty seat")
}

func findSeatId(line string) int {
	row := findPosition(line[:len(line)-3])
	col := findPosition(line[7:])
	return id(row, col)
}

func id(row int, col int) int {
	return row*8 + col
}

func findPosition(directions string) int {
	factor := int(math.Pow(2, float64(len(directions))))
	position := 0
	for _, d := range directions {
		if d == 'B' || d == 'R' {
			position += factor / 2
		}
		factor /= 2
	}
	return position
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
