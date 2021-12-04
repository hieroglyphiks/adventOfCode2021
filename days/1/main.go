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

	lines := input.lines()
	scans := []*scan{}
	for _, l := range lines {
		s, err := NewScan(l)
		if err != nil {
			log.Fatalf("failed to parse command: %s", l)
		}

		scans = append(scans, s)
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing scans...")
	increases := 0

	for idx, s := range scans {
		if idx < 1 {
			continue
		}

		if prevScan := scans[idx-1]; s.depth > prevScan.depth {
			increases++
		}
	}

	log.Println("Done")
	log.Printf("Depth Increases: %d", increases)

	// PART 1
	//=======
	log.Println("Part 2")
	log.Println("Processing scans...")
	increases = 0

	for idx, s := range scans {
		if idx-3 < 0 {
			continue
		}

		if prevScan := scans[idx-3]; s.depth > prevScan.depth {
			increases++
		}
	}

	log.Println("Done")
	log.Printf("Depth Increases: %d", increases)
}
