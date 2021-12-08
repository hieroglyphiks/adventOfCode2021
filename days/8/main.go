package main

import (
	"flag"
	"log"
	"math/bits"
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
	observations := []*observation{}
	for _, l := range lines {
		o, err := NewObservation(l)
		if err != nil {
			log.Fatalf("failed to parse observation: %s", err.Error())
		}

		observations = append(observations, o)
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing observations...")

	uniques := 0
	for _, observation := range observations {
		for _, output := range observation.output {
			switch len(output) {
			case 2, 3, 4, 7:
				uniques++
			default:
			}
		}
	}

	log.Println("Done")
	log.Printf("Uniques: %d", uniques)

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Processing observations...")

	total := 0
	for _, o := range observations {
		total += sum(o)
	}

	log.Println("Done")
	log.Printf("Total: %d", total)
}

func chars(field string) int {
	var res int
	for i := range field {
		res |= 1 << (field[i] - 'a')
	}
	return res
}

func sum(obs *observation) int {
	var matchChars [10]int
	for _, field := range obs.signals {
		switch len(field) {
		case 2:
			matchChars[1] = chars(field)
		case 3:
			matchChars[7] = chars(field)
		case 4:
			matchChars[4] = chars(field)
		case 7:
			matchChars[8] = chars(field)
		}
	}

	var sum int
	for _, field := range obs.output {
		contents := chars(field)
		sum *= 10
		switch len(field) {
		case 2:
			sum += 1
		case 3:
			sum += 7
		case 4:
			sum += 4
		case 5:
			switch {
			case contents&matchChars[1] == matchChars[1]:
				sum += 3
			case bits.OnesCount(uint(contents&matchChars[4])) == 3:
				sum += 5
			default:
				sum += 2
			}
		case 6:
			switch {
			case contents&matchChars[1] != matchChars[1]:
				sum += 6
			case contents&matchChars[4] == matchChars[4]:
				sum += 9
			default:
			}
		case 7:
			sum += 8
		}
	}
	return sum
}
