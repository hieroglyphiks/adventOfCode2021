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
	bingoOrder, err := NewOrder(lines[0])
	if err != nil {
		log.Fatalf("failed to parse bingo order: %s", err.Error())
	}

	boards := []*board{}
	p2boards := []*board{}
	for idx := 2; idx < len(lines); idx += 6 {
		board, err := NewBoard(lines[idx:idx+5], idx)
		if err != nil {
			log.Fatalf("failed to parse board: %s", err.Error())
		}
		boards = append(boards, board)

		p2board, err := NewBoard(lines[idx:idx+5], idx)
		if err != nil {
			log.Fatalf("failed to parse board: %s", err.Error())
		}
		p2boards = append(p2boards, p2board)
	}

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Filling order...")
	score := -1

	for _, o := range bingoOrder {
		for _, b := range boards {
			b.fill(o)

			if b.winner() {
				log.Println(b.id)
				score = b.unmarkedSum() * o
				break
			}
		}
		if score != -1 {
			break
		}
	}

	log.Printf("Final Score: %d\n", score)
	log.Println("Done")

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Filling order...")

	winners := map[int]struct{}{}

	var lastWinner *board
	lastWinnerScore := 0

	for _, o := range bingoOrder {
		for _, b := range p2boards {
			if _, ok := winners[b.id]; ok {
				continue
			}

			b.fill(o)
			if b.winner() {
				winners[b.id] = struct{}{}
				lastWinner = b
				lastWinnerScore = b.unmarkedSum() * o
			}
		}
	}

	log.Println(lastWinner.id)
	log.Printf("Final Score: %d\n", lastWinnerScore)
	log.Println("Done")
}
