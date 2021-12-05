package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// RAW INPUT READ
//===============
type input struct {
	body string
}

func NewInput(path string) (*input, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// clean
	cleanbody := strings.Trim(string(body), "\n")

	i := &input{
		body: string(cleanbody),
	}

	return i, nil
}

func (i *input) lines() []string {
	return strings.Split(i.body, "\n")
}

// PROBLEM SPECIFIC PARSE
type coordinate struct {
	x int
	y int
}

type line struct {
	start *coordinate
	end   *coordinate
}

func NewLine(inputline string) (*line, error) {
	l := &line{}

	coordinates := strings.Split(inputline, " -> ")
	if len(coordinates) != 2 {
		return nil, fmt.Errorf("invalid coordinates length: %d", len(coordinates))
	}

	// start
	startCoordinates := strings.Split(coordinates[0], ",")
	x1, err := strconv.Atoi(startCoordinates[0])
	if err != nil {
		return nil, fmt.Errorf("failed to convert start x to int: %w", err)
	}

	y1, err := strconv.Atoi(startCoordinates[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert start y to int: %w", err)
	}
	l.start = &coordinate{
		x: x1,
		y: y1,
	}

	// end
	endCoordinates := strings.Split(coordinates[1], ",")
	x2, err := strconv.Atoi(endCoordinates[0])
	if err != nil {
		return nil, fmt.Errorf("failed to convert end x to int: %w", err)
	}

	y2, err := strconv.Atoi(endCoordinates[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert end y to int: %w", err)
	}
	l.end = &coordinate{
		x: x2,
		y: y2,
	}

	return l, nil
}

type grid struct {
	// x->y->count
	coordinates map[int]map[int]int
}

func NewStraightOnlyGrid(lines []*line) (*grid, error) {
	g := &grid{
		coordinates: make(map[int]map[int]int, 0),
	}

	for _, l := range lines {
		// check if line is horizontal or vertical
		if !(l.start.x == l.end.x || l.start.y == l.end.y) {
			continue
		}

		if l.start.x == l.end.x {
			starty := int(math.Min(float64(l.start.y), float64(l.end.y)))
			endy := int(math.Max(float64(l.start.y), float64(l.end.y)))

			if _, ok := g.coordinates[l.start.x]; !ok {
				g.coordinates[l.start.x] = make(map[int]int, 0)
			}

			for y := starty; y <= endy; y++ {
				g.coordinates[l.start.x][y]++
			}
		}

		if l.start.y == l.end.y {
			startx := int(math.Min(float64(l.start.x), float64(l.end.x)))
			endx := int(math.Max(float64(l.start.x), float64(l.end.x)))

			for x := startx; x <= endx; x++ {
				if _, ok := g.coordinates[x]; !ok {
					g.coordinates[x] = make(map[int]int, 0)
				}

				g.coordinates[x][l.start.y]++
			}
		}
	}

	return g, nil
}

func NewGrid(lines []*line) (*grid, error) {
	g := &grid{
		coordinates: make(map[int]map[int]int, 0),
	}

	for _, l := range lines {
		// vertical
		if l.start.x == l.end.x {
			starty := int(math.Min(float64(l.start.y), float64(l.end.y)))
			endy := int(math.Max(float64(l.start.y), float64(l.end.y)))

			if _, ok := g.coordinates[l.start.x]; !ok {
				g.coordinates[l.start.x] = make(map[int]int, 0)
			}

			for y := starty; y <= endy; y++ {
				g.coordinates[l.start.x][y]++
			}
			continue
		}

		// horizontal
		if l.start.y == l.end.y {
			startx := int(math.Min(float64(l.start.x), float64(l.end.x)))
			endx := int(math.Max(float64(l.start.x), float64(l.end.x)))

			for x := startx; x <= endx; x++ {
				if _, ok := g.coordinates[x]; !ok {
					g.coordinates[x] = make(map[int]int, 0)
				}

				g.coordinates[x][l.start.y]++
			}
			continue
		}

		// diagonal
		// order by smaller x
		lowerxcoord := l.start
		higherxcoord := l.end

		if l.end.x < l.start.x {
			lowerxcoord = l.end
			higherxcoord = l.start
		}

		if lowerxcoord.y < higherxcoord.y {
			for x := lowerxcoord.x; x <= higherxcoord.x; x++ {
				if _, ok := g.coordinates[x]; !ok {
					g.coordinates[x] = make(map[int]int, 0)
				}

				y := lowerxcoord.y + (x - lowerxcoord.x)
				g.coordinates[x][y]++
			}
		}

		if lowerxcoord.y > higherxcoord.y {
			for x := lowerxcoord.x; x <= higherxcoord.x; x++ {
				if _, ok := g.coordinates[x]; !ok {
					g.coordinates[x] = make(map[int]int, 0)
				}

				y := lowerxcoord.y - (x - lowerxcoord.x)
				g.coordinates[x][y]++
			}
		}
	}

	return g, nil
}

func (g *grid) AtLeastTwoOverlap() int {
	sum := 0
	for _, ys := range g.coordinates {
		for _, overlappinglines := range ys {
			if overlappinglines > 1 {
				sum++
			}
		}
	}

	return sum
}
