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
	diagnostics := []*diagnostic{}
	for _, l := range lines {
		d, err := NewDiagnostic(l)
		if err != nil {
			log.Fatalf("failed to parse diagnostic: %s", l)
		}

		diagnostics = append(diagnostics, d)
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Getting length...")

	maxLen := 0
	for _, d := range diagnostics {
		if d.length > maxLen {
			maxLen = d.length
		}
	}

	log.Println("Done.")
	log.Printf("MaxLen: %d\n", maxLen)
	log.Println("Processing diagnostics...")
	gammaRate := 0
	epsilonRate := 0

	for bitIdx := 0; bitIdx < maxLen; bitIdx++ {
		mcb := mostCommonBit(diagnostics, bitIdx)

		if mcb == 1 {
			gammaRate += int(math.Pow(float64(2), float64(bitIdx)))
		} else {
			epsilonRate += int(math.Pow(float64(2), float64(bitIdx)))
		}

	}

	log.Println("Done")
	log.Printf("GammaRate: %d, EpsilonRate: %d, Solution: %d\n", gammaRate, epsilonRate, gammaRate*epsilonRate)

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Finding Oxygen Generator Rating...")

	diagnosticSet := diagnostics
	bitIdx := maxLen - 1

	for len(diagnosticSet) > 1 {
		filteredSet := []*diagnostic{}
		mostCommon := mostCommonBit(diagnosticSet, bitIdx)

		for _, d := range diagnosticSet {
			diagValue := (d.binary & int(math.Pow(float64(2), float64(bitIdx)))) >> bitIdx

			if diagValue == mostCommon {
				filteredSet = append(filteredSet, d)
			}
		}

		diagnosticSet = filteredSet
		bitIdx--
	}

	oxygenRating := diagnosticSet[0].binary
	log.Println("Done.")

	log.Println("Finding C02 Scrubber Rating...")
	diagnosticSet = diagnostics
	bitIdx = maxLen - 1

	for len(diagnosticSet) > 1 {
		filteredSet := []*diagnostic{}
		leastCommon := mostCommonBit(diagnosticSet, bitIdx) ^ 1

		for _, d := range diagnosticSet {
			diagValue := (d.binary & int(math.Pow(float64(2), float64(bitIdx)))) >> bitIdx

			if diagValue == leastCommon {
				filteredSet = append(filteredSet, d)
			}
		}

		diagnosticSet = filteredSet
		bitIdx--
	}

	c02Rating := diagnosticSet[0].binary
	log.Println("Done.")
	log.Printf("Oxyen Rating: %d, c02 Rating: %d, Solution: %d\n", oxygenRating, c02Rating, oxygenRating*c02Rating)
}

func mostCommonBit(ds []*diagnostic, bitIdx int) int {
	bitSum := 0
	for _, d := range ds {
		bitSum += (d.binary & int(math.Pow(float64(2), float64(bitIdx)))) >> bitIdx
	}

	if float64(bitSum)/float64(len(ds)) >= 0.5 {
		return 1
	} else {
		return 0
	}
}
