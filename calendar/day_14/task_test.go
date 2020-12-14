package main

import (
	"strings"
	"testing"
)

const input = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

const input2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`


var lines = strings.Split(input, "\n")
var lines2 = strings.Split(input2, "\n")

const expectedResultPart1 = int64(165)
const expectedResultPart2 = int64(208)

func TestPart1Example(t *testing.T) {
	result, err := part1(lines)

	validateResult(t, err, result, expectedResultPart1)
}

func TestPart2Example(t *testing.T) {
	result, err := part2(lines2)

	validateResult(t, err, result, expectedResultPart2)
}

func validateResult(t *testing.T, err error, result int64, expected int64) {
	if err != nil {
		t.Fatalf("Unexpected Error: %s", err)
	}

	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}
