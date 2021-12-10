package main

import (
	"flag"
	"log"
	"sort"
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

	// PART 1
	//=======
	log.Println("Part 1")
	log.Println("Processing lines ...")

	scoreSum := 0
	for _, l := range lines {
		fi, _ := firstIllegal(l)
		if fi != "" {
			scoreSum += score(fi)
		}
	}

	log.Println("Done")
	log.Printf("score: %d\n", scoreSum)

	// PART 2
	//=======
	log.Println("Part 2")
	log.Println("Processing lines ...")
	scores := []int{}
	for _, l := range lines {
		_, c := firstIllegal(l)
		if c != "" {
			scores = append(scores, incompleteScore(c))
		}
	}

	sort.Ints(scores)

	log.Println("Done")
	log.Printf("score: %d\n", scores[int(float64(len(scores))/float64(2))])
}

type char struct {
	val   string
	idx   []int
	count int
}

func firstIllegal(line string) (string, string) {
	openCounts := map[string]*char{}
	for idx := 0; idx < len(line); idx++ {
		c := string(line[idx])
		switch c {
		case "(", "[", "{", "<":
			openChar, ok := openCounts[c]
			if !ok {
				openCounts[c] = &char{
					val:   c,
					idx:   []int{idx},
					count: 1,
				}
			} else {
				openChar.idx = append(openChar.idx, idx)
				openChar.count++
			}
		case ")", "]", "}", ">":
			opposite := closeOpposite(c)

			open, ok := openCounts[opposite]
			if ok {
				if open.count < 1 {
					return c, ""
				}

				subIllegal, incomplete := firstIllegal(line[open.idx[len(open.idx)-1]+1 : idx])
				if subIllegal != "" {
					return subIllegal, ""
				} else if incomplete != "" {
					return c, ""
				}
			}

			open.count--
			open.idx = open.idx[:len(open.idx)-1]
		}
	}

	idxs := map[int]string{}
	keys := []int{}
	for _, v := range openCounts {
		for _, i := range v.idx {
			idxs[i] = v.val
			keys = append(keys, i)
		}
	}

	sort.Ints(keys)
	completion := ""
	for i := len(keys) - 1; i >= 0; i-- {
		completion += openOpposite(idxs[keys[i]])
	}

	return "", completion
}

func closeOpposite(c string) string {
	switch c {
	case ")":
		return "("
	case "]":
		return "["
	case "}":
		return "{"
	case ">":
		return "<"
	}

	return ""
}

func openOpposite(c string) string {
	switch c {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}

	return ""
}

func score(c string) int {
	switch c {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}

	return 0
}

func incompleteScore(c string) int {
	score := 0
	for i := 0; i < len(c); i++ {
		v := 0
		switch string(c[i]) {
		case ")":
			v = 1
		case "]":
			v = 2
		case "}":
			v = 3
		case ">":
			v = 4
		}

		score = (score * 5) + v
	}

	return score
}
