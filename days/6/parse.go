package main

import (
	"fmt"
	"io/ioutil"
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
type days map[int]int

func NewDays(line string) (days, error) {
	d := days{}

	numbers := strings.Split(line, ",")

	// prepopulate days
	for x := 0; x <= 8; x++ {
		d[x] = 0
	}

	for _, n := range numbers {
		number, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number: %s", n)
		}

		d[number] = d[number] + 1
	}

	return d, nil
}
