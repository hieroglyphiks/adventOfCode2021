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
type diagnostic struct {
	binary int
	length int
}

func NewDiagnostic(line string) (*diagnostic, error) {
	d := &diagnostic{}
	d.length = len(line)

	binary, err := strconv.ParseInt(line, 2, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse binary: %w", err)
	}
	d.binary = int(binary)

	return d, nil
}
