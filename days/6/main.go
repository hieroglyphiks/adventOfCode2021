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
	fishDays, err := NewDays(lines[0])
	if err != nil {
		log.Fatalf("failed to parse days: %s", err.Error())
	}

	fishDays2, err := NewDays(lines[0])
	if err != nil {
		log.Fatalf("failed to parse days: %s", err.Error())
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing days...")

	for day := 0; day < 80; day++ {
		birthing := fishDays[0]

		for fd := 0; fd < 8; fd++ {
			fishDays[fd] = fishDays[fd+1]
		}

		fishDays[6] = fishDays[6] + birthing
		fishDays[8] = birthing
	}

	sum := 0
	for _, count := range fishDays {
		sum += count
	}

	log.Println("Done")
	log.Printf("Fish Sum: %d\n", sum)

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Processing days...")

	for day := 0; day < 256; day++ {
		birthing := fishDays2[0]

		for fd := 0; fd < 8; fd++ {
			fishDays2[fd] = fishDays2[fd+1]
		}

		fishDays2[6] = fishDays2[6] + birthing
		fishDays2[8] = birthing
	}

	sum2 := 0
	for _, count := range fishDays2 {
		sum2 += count
	}

	log.Println("Done")
	log.Printf("Fish Sum: %d\n", sum2)
}
