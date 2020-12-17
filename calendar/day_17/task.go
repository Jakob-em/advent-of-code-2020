package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type position struct {
	x int
	y int
	z int
	w int
}

type Conway struct {
	fourDimensions bool
	field          map[position]bool
}

type positionVisitor func(pos position)

func NewConway(lines []string, fourDimensions bool) Conway {
	field := map[position]bool{}
	for y, line := range lines {
		for x, c := range line {
			field[position{
				x: x,
				y: y,
			}] = c == '#'
		}
	}
	return Conway{
		fourDimensions: fourDimensions,
		field:          field,
	}
}

func (c *Conway) simulate(turns int) {
	for i := 0; i < turns; i++ {
		newField := map[position]bool{}
		for p := range c.field {
			c.processNeighbors(p, true, func(pos position) {
				state := c.field[pos]
				activeNeighbors := c.countActiveNeighbors(pos)
				if state {
					newField[pos] = activeNeighbors == 3 || activeNeighbors == 2
				}
				if !state && activeNeighbors == 3 {
					newField[pos] = true
				}
			})
		}
		c.field = newField
	}
}

func (c *Conway) countActiveFields() int {
	count := 0
	for _, s := range c.field {
		if s {
			count++
		}
	}
	return count
}

func (c *Conway) countActiveNeighbors(pos position) int {
	count := 0
	c.processNeighbors(pos, false, func(pos position) {
		if c.field[pos] {
			count++
		}
	})
	return count
}

func (c *Conway) processNeighbors(pos position, includeOwn bool, visit positionVisitor) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				excluded := x == 0 && y == 0 && z == 0 && !includeOwn

				if c.fourDimensions {
					for w := -1; w <= 1; w++ {
						if excluded && w == 0 {
							continue
						}
						visit(position{
							x: pos.x + x,
							y: pos.y + y,
							z: pos.z + z,
							w: pos.w + w,
						})
					}
				} else {
					if excluded {
						continue
					}
					visit(position{
						x: pos.x + x,
						y: pos.y + y,
						z: pos.z + z,
					})
				}
			}
		}
	}
}

func part1(lines []string) (int, error) {
	conway := NewConway(lines, false)
	conway.simulate(6)
	return conway.countActiveFields(), nil
}

func part2(lines []string) (int, error) {
	conway := NewConway(lines, true)
	conway.simulate(6)
	return conway.countActiveFields(), nil
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
