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

type scan struct {
	depth int
}

func NewScan(line string) (*scan, error) {
	s := &scan{}

	depth, err := strconv.Atoi(line)
	if err != nil {
		return nil, fmt.Errorf("failed to get parse depth: %w", err)
	}
	s.depth = depth

	return s, nil
}
