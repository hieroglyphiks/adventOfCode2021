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
type direction string

const (
	Forward direction = "forward"
	Down    direction = "down"
	Up      direction = "up"
)

type command struct {
	dir  direction
	dist int
}

func NewCommand(line string) (*command, error) {
	parts := strings.Split(strings.TrimSpace(line), " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("not a 2 part line: [%s]", line)
	}

	c := &command{}
	switch direction(parts[0]) {
	case Forward, Down, Up:
		c.dir = direction(parts[0])
	default:
		return nil, fmt.Errorf("invalid direction: [%s]", parts[0])
	}

	dist, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to get parse distance: %w", err)
	}
	c.dist = dist

	return c, nil
}
