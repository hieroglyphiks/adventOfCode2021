package main

import (
	"flag"
	"log"
	"math"
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

	lines := input.lines()
	crabPositions, err := NewPositions(lines[0])
	if err != nil {
		log.Fatalf("failed to parse days: %s", err.Error())
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing days...")

	min := math.MaxInt64
	for p := 0; p <= crabPositions.max(); p++ {
		dif := crabPositions.difference(p)
		if dif < min {
			min = dif
		}
	}

	log.Println("Done")
	log.Printf("Difference: %d", min)

	// PART 1
	//=======
	log.Println("Part 2")
	log.Println("Processing days...")

	min2 := math.MaxInt64
	for p := 0; p <= crabPositions.max(); p++ {
		dif := crabPositions.gausianDifference(p)
		if dif < min2 {
			min2 = dif
		}
	}

	log.Println("Done")
	log.Printf("Difference: %d", min2)
}
