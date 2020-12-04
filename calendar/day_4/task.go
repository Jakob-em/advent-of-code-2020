package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type validationFunction func(map[string]string) bool

var allPassportFields = []string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid",
}

var hairColorPattern = regexp.MustCompile(`^#(?:[0-9a-f]{6})$`)
var passportIdPattern = regexp.MustCompile(`^[0-9]{9}$`)
var allowedEyeColors = []string{
	"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
}

func part1(lines []string) (int, error) {
	return countValidPassportsInList(lines, arePassPortKeysValid), nil
}

func part2(lines []string) (int, error) {
	return countValidPassportsInList(lines, arePassPortValuesAndKeysValid), nil
}

func countValidPassportsInList(lines []string, validate validationFunction) int {
	validPassports := 0
	var currentPassPortFields = map[string]string{}

	for _, line := range lines {
		if line != "" {
			parsePassportLine(line, currentPassPortFields)
		} else {
			if validate(currentPassPortFields) {
				validPassports++
			}
			currentPassPortFields = map[string]string{}
		}
	}
	if validate(currentPassPortFields) {
		validPassports++
	}

	return validPassports
}

func arePassPortKeysValid(passportFields map[string]string) bool {
	_, containsCid := passportFields["cid"]
	isOnlyCidMissing := len(passportFields) == len(allPassportFields)-1 && !containsCid
	areAllFieldsPresent := len(passportFields) == len(allPassportFields)
	return areAllFieldsPresent || isOnlyCidMissing
}

func arePassPortValuesAndKeysValid(passportFields map[string]string) bool {
	if !arePassPortKeysValid(passportFields) {
		return false
	}

	if !isMapEntryInRange(passportFields, "byr", 1920, 2002) {
		return false
	}
	if !isMapEntryInRange(passportFields, "iyr", 2010, 2020) {
		return false
	}
	if !isMapEntryInRange(passportFields, "eyr", 2020, 2030) {
		return false
	}

	hgt := passportFields["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		if !validateHeightValue(hgt, "cm", 150, 193) {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		if !validateHeightValue(hgt, "in", 59, 76) {
			return false
		}
	} else {
		return false
	}

	hcl := passportFields["hcl"]
	if !hairColorPattern.MatchString(hcl) {
		return false
	}

	ecl := passportFields["ecl"]
	if !isValueInWhitelist(ecl, allowedEyeColors) {
		return false
	}

	pid := passportFields["pid"]
	if !passportIdPattern.MatchString(pid) {
		return false
	}

	return true
}

func validateHeightValue(hgt string, suffix string, min int, max int) bool {
	intValue, err := strconv.Atoi(strings.Split(hgt, suffix)[0])
	if err != nil {
		return false
	}
	return intValue >= min && intValue <= max
}

func isMapEntryInRange(m map[string]string, key string, min int, max int) bool {
	value := m[key]
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return intValue >= min && intValue <= max
}
func isValueInWhitelist(value string, whitelist []string) bool {
	isValid := false
	for _, allowedValue := range whitelist {
		if allowedValue == value {
			isValid = true
			break
		}
	}
	return isValid
}

func parsePassportLine(line string, currentPassPortFields map[string]string) {
	groups := strings.Split(line, " ")
	for _, group := range groups {
		kv := strings.Split(group, ":")
		currentPassPortFields[kv[0]] = kv[1]
	}
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
