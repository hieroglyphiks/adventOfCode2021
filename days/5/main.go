package main

import (
	"flag"
	"log"
)

func main() {
	// FLAGS
	//======
	var inputFilepath string

	flag.StringVar(&inputFilepath, "i", "input.txt", "input filepath")
	flag.Parse()

	// PARSE INPUT
	//============
	input, err := NewInput(inputFilepath)
	if err != nil {
		log.Fatalf("failed to parse input: %s", err.Error())
	}

	inputlines := input.lines()
	lines := []*line{}
	for _, l := range inputlines {
		l, err := NewLine(l)
		if err != nil {
			log.Fatalf("failed to parse line: %+v", l)
		}

		lines = append(lines, l)
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing lines...")

	grid, err := NewStraightOnlyGrid(lines)
	if err != nil {
		log.Fatalf("failed to create straight only grid: %s", err.Error())
	}

	points := grid.AtLeastTwoOverlap()
	log.Println("Done")
	log.Printf("Points with at least 2 overlapping lines: %d\n", points)

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Processing lines...")

	grid2, err := NewGrid(lines)
	if err != nil {
		log.Fatalf("failed to create grid: %s", err.Error())
	}

	points2 := grid2.AtLeastTwoOverlap()
	log.Println("Done")
	log.Printf("Points with at least 2 overlapping lines: %d\n", points2)
}
