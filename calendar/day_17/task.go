package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/Jakob-em/advent-of-code-2020/utils"
)

type position struct {
	x int
	y int
	z int
	w int
}

type conway struct {
	fourDimensions bool
	field          map[position]bool
}

type positionVisitor func(pos position)

func newConway(lines []string, fourDimensions bool) conway {
	field := map[position]bool{}
	for y, line := range lines {
		for x, c := range line {
			field[position{
				x: x,
				y: y,
			}] = c == '#'
		}
	}
	return conway{
		fourDimensions: fourDimensions,
		field:          field,
	}
}

func (c *conway) simulate(turns int) {
	for i := 0; i < turns; i++ {
		newActiveCells := make(chan position, 1000)
		var wg sync.WaitGroup

		for p := range c.field {
			wg.Add(1)
			go func(p position) {
				defer wg.Done()
				c.processNeighbors(p, func(pos position) {
					state := c.field[pos]
					activeNeighbors := c.countActiveNeighbors(pos)
					if state && (activeNeighbors == 3 || activeNeighbors == 2) {
						newActiveCells <- pos
					} else if !state && activeNeighbors == 3 {
						newActiveCells <- pos
					}
				})
			}(p)
		}

		newFieldChan := make(chan map[position]bool)

		go func() {
			newField := map[position]bool{}
			for p := range newActiveCells {
				newField[p] = true
			}
			newFieldChan <- newField
		}()

		wg.Wait()
		close(newActiveCells)
		c.field = <-newFieldChan
	}
}

func (c *conway) countActiveFields() int {
	count := 0
	for _, s := range c.field {
		if s {
			count++
		}
	}
	return count
}

func (c *conway) countActiveNeighbors(pos position) int {
	count := 0
	c.processNeighbors(pos, func(p position) {
		if pos != p && c.field[p] {
			count++
		}
	})
	return count
}

func (c *conway) processNeighbors(pos position, visit positionVisitor) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if c.fourDimensions {
					for w := -1; w <= 1; w++ {
						visit(position{
							x: pos.x + x,
							y: pos.y + y,
							z: pos.z + z,
							w: pos.w + w,
						})
					}
				} else {
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
	c := newConway(lines, false)
	c.simulate(6)
	return c.countActiveFields(), nil
}

func part2(lines []string) (int, error) {
	c := newConway(lines, true)
	c.simulate(6)
	return c.countActiveFields(), nil
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
