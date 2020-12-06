package main

import (
	"strings"
	"testing"
)

const input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

var lines = strings.Split(input, "\n")

const expectedResultPart1 = 2
const expectedResultPart2 = 1

func TestPart1Example(t *testing.T) {
	result, err := part1(lines)

	validateResult(t, err, result, expectedResultPart1)
}

func TestPart2Example(t *testing.T) {
	result, err := part2(lines)

	validateResult(t, err, result, expectedResultPart2)
}

func validateResult(t *testing.T, err error, result int, expected int) {
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err)
	}

	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}
