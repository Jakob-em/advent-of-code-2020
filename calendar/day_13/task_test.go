package main

import (
	"strings"
	"testing"
)

const input = `939
7,13,x,x,59,x,31,19`

var lines = strings.Split(input, "\n")

const expectedResultPart1 = 295
const expectedResultPart2 = 1068781

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
