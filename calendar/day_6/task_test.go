package main

import (
	"strings"
	"testing"
)

const input = `abc

a
b
c

ab
ac

a
a
a
a

b`

var lines = strings.Split(input, "\n\n")

const expectedResultPart1 = 11
const expectedResultPart2 = 6

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
