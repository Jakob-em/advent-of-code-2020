package main

import (
	"strings"
	"testing"
)

const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

var lines = strings.Split(input, "\n")

const expectedResultPart1 = 127
const expectedResultPart2 = 62

func TestPart1Example(t *testing.T) {
	result, err := part1(lines, 5)

	validateResult(t, err, result, expectedResultPart1)
}

func TestPart2Example(t *testing.T) {
	result, err := part2(lines, 5)

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
