package main

import (
	"testing"
)

var testData = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

const expectedResultPart1 = 7

const expectedResultPart2 = 336

func TestPart1Example(t *testing.T) {
	result, err := Part1(testData)

	validateResult(t, err, result, expectedResultPart1)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2(testData)

	validateResult(t, err, result, expectedResultPart2)
}

func validateResult(t *testing.T, err error, result int, expected int) {
	if err != nil {
		t.Fatalf("Unexcpected Error: %s", err)
	}

	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}
