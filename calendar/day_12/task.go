package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type Point struct {
	x int
	y int
}

var directionMapping = map[uint8]int{'N': 0, 'E': 1, 'S': 2, 'W': 3}

func part1(lines []string) (int, error) {

	shipPos := Point{
		x: 0,
		y: 0,
	}
	shipDir := 1

	for _, line := range lines {
		action := line[0]
		arg, _ := strconv.Atoi(line[1:])

		if action == 'N' || action == 'S' || action == 'E' || action == 'W' {
			shipPos.movePoint(directionMapping[action], arg)
		} else if action == 'L' {
			shipDir += -(arg / 90) + 4
			shipDir %= 4
		} else if action == 'R' {
			shipDir += arg / 90
			shipDir %= 4
		} else if action == 'F' {
			shipPos.movePoint(shipDir, arg)
		}
	}

	return Abs(shipPos.x) + Abs(shipPos.y), nil
}

func (pos *Point) movePoint(direction int, amount int) {
	if direction == 0 {
		pos.y += amount
	} else if direction == 1 {
		pos.x += amount
	} else if direction == 2 {
		pos.y -= amount
	} else if direction == 3 {
		pos.x -= amount
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(lines []string) (int, error) {

	shipPos := Point{
		x: 0,
		y: 0,
	}
	waypoint := Point{
		x: 10,
		y: 1,
	}

	for _, line := range lines {
		action := line[0]
		arg, _ := strconv.Atoi(line[1:])

		if action == 'N' || action == 'S' || action == 'E' || action == 'W' {
			waypoint.movePoint(directionMapping[action], arg)
		} else if action == 'L' {
			waypoint.rotate(((-arg / 90) + 4) % 4)
		} else if action == 'R' {
			waypoint.rotate((arg / 90) % 4)
		} else if action == 'F' {
			shipPos.x += arg * waypoint.x
			shipPos.y += arg * waypoint.y
		}
	}

	return Abs(shipPos.x) + Abs(shipPos.y), nil
}

func (pos *Point) rotate(amount int) {
	if amount == 1 {
		old := *pos
		pos.x = old.y
		pos.y = -old.x
	} else if amount == 2 {
		pos.x = - pos.x
		pos.y = - pos.y
	} else if amount == 3 {
		old := *pos
		pos.x = -old.y
		pos.y = old.x
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
