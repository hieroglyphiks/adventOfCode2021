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
//=======================
type positions map[int]int

func NewPositions(line string) (positions, error) {
	p := positions{}

	numbers := strings.Split(line, ",")

	// prepopulate days
	for _, n := range numbers {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number: %s", n)
		}

		p[number] = p[number] + 1
	}

	return p, nil
}

func (p positions) max() int {
	max := 0
	for pos := range p {
		if pos > max {
			max = pos
		}
	}
	return max
}

func (p positions) difference(goal int) int {
	difference := 0

	for position, count := range p {
		difference += int(math.Abs(float64(position-goal))) * count
	}

	return difference
}

func (p positions) gausianDifference(goal int) int {
	difference := 0

	for position, count := range p {
		distance := int(math.Abs(float64(position - goal)))
		gaussian := int((float64(distance) / float64(2)) * float64(1+distance))
		difference += gaussian * count
	}

	return difference
}

func gausian(position int, goal int) int {
	distance := int(math.Abs(float64(position - goal)))
	return int((float64(distance) / float64(2)) * float64(1+distance))
}
