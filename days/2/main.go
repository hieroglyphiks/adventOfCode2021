package main

import (
	"flag"
	"log"
)

func main() {
	// FLAGS
	//======
	var inputFilepath string

	flag.StringVar(&inputFilepath, "i", "", "input filepath")
	flag.Parse()

	// PARSE INPUT
	//============
	input, err := NewInput(inputFilepath)
	if err != nil {
		log.Fatalf("failed to parse input: %s", err.Error())
	}

	lines := input.lines()
	commands := []*command{}
	for _, l := range lines {
		c, err := NewCommand(l)
		if err != nil {
			log.Fatalf("failed to parse command: %s", l)
		}

		commands = append(commands, c)
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing commands...")
	xpos := 0
	depth := 0

	for _, c := range commands {
		switch c.dir {
		case Forward:
			xpos += c.dist
		case Up:
			depth -= c.dist
		case Down:
			depth += c.dist
		}
	}

	log.Println("Done")
	log.Printf("Final Positions -> Xpos: %d, Depth: %d", xpos, depth)
	log.Printf("Solution: %d\n", xpos*depth)

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Processing commands...")
	xpos = 0
	depth = 0
	aim := 0

	for _, c := range commands {
		switch c.dir {
		case Forward:
			xpos += c.dist
			depth += c.dist * aim
		case Up:
			aim -= c.dist
		case Down:
			aim += c.dist
		}
	}

	log.Println("Done")
	log.Printf("Final Positions -> Xpos: %d, Depth: %d", xpos, depth)
	log.Printf("Solution: %d\n", xpos*depth)
}
