package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type bagDescription struct {
	color string
	count int
}

type combination struct {
	parent       string
	destinations []bagDescription
}

type bagRelations = map[string][]bagDescription

func part1(lines []string) (int, error) {

	relations := extractBagRelations(lines)

	parentRelations := map[string][]string{}
	for parent, destinations := range relations {
		for _, dest := range destinations {
			parentRelations[dest.color] = append(parentRelations[dest.color], parent)
		}
	}
	possibleParents := map[string]bool{}
	findPossibleParents(parentRelations, "shiny gold", possibleParents)
	return len(possibleParents), nil
}

func findPossibleParents(parentRelations map[string][]string, search string, possibleParents map[string]bool) {
	parents, contained := parentRelations[search]
	if !contained {
		return
	}

	for _, parent := range parents {
		possibleParents[parent] = true
		findPossibleParents(parentRelations, parent, possibleParents)
	}
}

func part2(lines []string) (int, error) {
	relations := extractBagRelations(lines)

	return countNeededBags(relations, "shiny gold"), nil
}

func countNeededBags(relations bagRelations, search string) int {
	_, contained := relations[search]
	if !contained {
		return 0
	}
	result := 0
	for _, d := range relations[search] {
		result += d.count + d.count*countNeededBags(relations, d.color)
	}
	return result
}

func extractBagRelations(lines []string) bagRelations {
	combinations := bagRelations{}

	for _, line := range lines {
		combination, _ := parseLine(line)
		for _, d := range combination.destinations {
			combinations[combination.parent] = append(combinations[combination.parent], d)
		}
	}
	return combinations
}

func parseLine(line string) (combination, error) {
	parentGroups := utils.ExtractGroups(`^(?P<src>.+) bag[s] contain (?P<dest>.+)\.$`, line)
	destinationTexts := strings.Split(parentGroups["dest"], ", ")

	var destinations []bagDescription
	for _, text := range destinationTexts {
		if strings.Contains(text, "no other bags") {
			break
		}
		destinationGroups := utils.ExtractGroups(`^(?P<count>\d+) (?P<dest>.+) bags?\.?$`, text)

		count, _ := strconv.Atoi(destinationGroups["count"])
		destinations = append(destinations, bagDescription{
			color: destinationGroups["dest"],
			count: count,
		})
	}

	return combination{parent: parentGroups["src"], destinations: destinations}, nil
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
