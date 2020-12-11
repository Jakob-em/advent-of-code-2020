package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

func part1(lines []string) (int, error) {
	grid := initGrid(lines)
	for {
		newGrid := CalculateRound(grid, 4, CountAdjacentOccupiedFieldsWithContent)
		if CompareGrids(newGrid, grid) {
			break
		}
		grid = newGrid
	}

	return countOccupiedSeats(grid), nil
}

func countOccupiedSeats(grid [][]rune) int {
	occupied := 0
	for _, rowContent := range grid {
		for _, c := range rowContent {
			if c == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func initGrid(lines []string) [][]rune {
	rows := len(lines)
	cols := len(lines[0])

	grid := CreateGrid(rows, cols)

	for row, line := range lines {
		for col, c := range line {
			grid[row][col] = c
		}
	}
	return grid
}

type occupiedFunction func(grid [][]rune, posX int, posY int) int

func CalculateRound(oldGrid [][]rune, occupiedLimit int, fun occupiedFunction) [][]rune {
	grid := CreateGrid(len(oldGrid), len(oldGrid[0]))

	for row, rowContent := range oldGrid {
		for col, c := range rowContent {

			if c == '.' {
				grid[row][col] = '.'

			} else {
				occupiedSeats := fun(oldGrid, col, row)
				if c == 'L' && occupiedSeats == 0 {
					grid[row][col] = '#'
				} else if c == '#' && occupiedSeats >= occupiedLimit {
					grid[row][col] = 'L'
				} else {
					grid[row][col] = oldGrid[row][col]
				}
			}

		}
	}
	return grid
}

func CompareGrids(grid1 [][]rune, grid2 [][]rune) bool {
	for row, rowContent := range grid1 {
		for col := range rowContent {
			if grid1[row][col] != grid2[row][col] {
				return false
			}
		}
	}
	return true
}

func CountAdjacentOccupiedFieldsWithContent(grid [][]rune, posX int, posY int) int {
	count := 0
	for row := posY - 1; row <= posY+1; row++ {
		for col := posX - 1; col <= posX+1; col++ {
			if !(posX == col && posY == row) && !isOutOfGrid(grid, col, row) {
				if grid[row][col] == '#' {
					count++
				}
			}
		}
	}
	return count
}

func CreateGrid(rows int, cols int) [][]rune {
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
	}
	return grid
}

func part2(lines []string) (int, error) {
	grid := initGrid(lines)
	for {
		newGrid := CalculateRound(grid, 5, CountOccupiedVisibleFields)
		if CompareGrids(newGrid, grid) {
			break
		}
		grid = newGrid
	}

	return countOccupiedSeats(grid), nil
}

type direction struct {
	dirX int
	dirY int
}

var directions = []direction{
	{
		dirX: -1,
		dirY: -1,
	},
	{
		dirX: -1,
		dirY: 0,
	}, {
		dirX: 0,
		dirY: -1,
	}, {
		dirX: 0,
		dirY: 1,
	},
	{
		dirX: 1,
		dirY: 0,
	}, {
		dirX: 1,
		dirY: 1,
	},
	{
		dirX: -1,
		dirY: 1,
	}, {
		dirX: 1,
		dirY: -1,
	},
}

func CountOccupiedVisibleFields(grid [][]rune, posX int, posY int) int {
	count := 0
	for _, d := range directions {
		if isSeatOccupiedInDirection(grid, posX, posY, d) {
			count++
		}
	}
	return count
}

func isSeatOccupiedInDirection(grid [][]rune, posX int, posY int, dir direction) bool {
	posX += dir.dirX
	posY += dir.dirY
	for !isOutOfGrid(grid, posX, posY) {
		c := grid[posY][posX]
		if c == '#' {
			return true
		} else if c == 'L' {
			return false
		}
		posX += dir.dirX
		posY += dir.dirY
	}
	return false
}

func isOutOfGrid(grid [][]rune, posX int, posY int) bool {
	if posY < 0 || posY >= len(grid) {
		return true
	}
	if posX < 0 || posX >= len(grid[0]) {
		return true
	}
	return false
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
