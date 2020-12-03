package main

import (
	"strings"
	"testing"
)

const input = `1721
979
366
299
675
1456`

var lines = strings.Split(input, "\n")

const expectedResultPart1 = 514579
const expectedResultPart2 = 241861950

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
		t.Fatalf("Unexcpected Error: %s", err)
	}

	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}
