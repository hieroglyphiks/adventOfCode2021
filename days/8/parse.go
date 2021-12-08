package main

import (
	"fmt"
	"io/ioutil"
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
type observation struct {
	signals []string
	output  []string
}

func NewObservation(line string) (*observation, error) {
	categories := strings.Split(strings.TrimSpace(line), "|")
	if len(categories) != 2 {
		return nil, fmt.Errorf("invalid line format")
	}

	signals := strings.Split(strings.TrimSpace(categories[0]), " ")
	output := strings.Split(strings.TrimSpace(categories[1]), " ")

	o := &observation{
		signals: signals,
		output:  output,
	}

	return o, nil
}
