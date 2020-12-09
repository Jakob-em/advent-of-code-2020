package main

import (
	"errors"
	"fmt"
	"github.com/Jakob-em/advent-of-code-2020/utils"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

func part1(lines []string, preambleSize int) (int, error) {
	nums, err := utils.ToIntSlice(lines)
	if err != nil {
		return 0, err
	}
	return findInvalidNumber(nums, preambleSize)
}

func findInvalidNumber(nums []int, preambleSize int) (int, error) {

	for i := preambleSize; i <= len(nums); i++ {
		lastNums := make([]int, preambleSize)
		copy(lastNums, nums[i-preambleSize:i])

		if !isSumInLines(lastNums, nums[i]) {
			return nums[i], nil
		}

	}

	return 0, errors.New("could not find invalid number")
}

func isSumInLines(nums []int, num int) bool {
	for i, num1 := range nums[:len(nums)-1] {
		for _, num2 := range nums[i+1:] {
			if num1+num2 == num {
				return true
			}
		}
	}
	return false
}

func part2(lines []string, preambleSize int) (int, error) {
	nums, err := utils.ToIntSlice(lines)
	if err != nil {
		return 0, err
	}

	invalid, _ := findInvalidNumber(nums, preambleSize)

	for i := range nums {
		low, high, err := findContinuousSum(i, invalid, nums)
		if err == nil {
			min := utils.Min(nums[low : high+1])
			max := utils.Max(nums[low : high+1])
			return min + max, nil
		}

	}

	return 0, errors.New("could not find encryption weakness")
}

func findContinuousSum(start int, invalid int, nums []int) (int, int, error) {
	sum := 0
	for end := start; end < len(nums); end++ {
		sum += nums[end]
		if sum == invalid {
			return start, end, nil
		}
	}
	return 0, 0, errors.New("could not find continuous sum")
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	inputFile := filepath.Join(path.Dir(filename), "input.txt")
	lines := utils.ReadLinesFromFile(inputFile, "\n")

	result, err := part1(lines, 25)
	if err != nil {
		log.Fatalf("Part 1: %s", err)
	}
	fmt.Printf("Result from Part 1: %d\n", result)

	result, err = part2(lines, 25)
	if err != nil {
		log.Fatalf("Part 2: %s", err)
	}
	fmt.Printf("Result from Part 2: %d\n", result)
}
