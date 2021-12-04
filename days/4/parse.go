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
type space struct {
	value  int
	filled bool
}

type board struct {
	id     int
	spaces [][]*space
	values map[int]*space
}

func NewBoard(lines []string, id int) (*board, error) {
	b := &board{
		id:     id,
		spaces: make([][]*space, 0),
		values: map[int]*space{},
	}

	if len(lines) != 5 {
		return nil, fmt.Errorf("invalid line count: %d", len(lines))
	}

	for _, l := range lines {
		spaceValues := strings.Split(strings.TrimSpace(l), " ")

		spaces := []*space{}
		for _, s := range spaceValues {
			if s == "" {
				continue
			}

			v, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("failed to convert space value: %w", err)
			}

			space := &space{
				value:  v,
				filled: false,
			}

			spaces = append(spaces, space)
			b.values[v] = space
		}

		b.spaces = append(b.spaces, spaces)
	}

	return b, nil
}

func (b *board) fill(v int) {
	if s, ok := b.values[v]; ok {
		s.filled = true
	}
}

func (b *board) unmarkedSum() int {
	sum := 0
	for _, row := range b.spaces {
		for _, col := range row {
			if !col.filled {
				sum += col.value
			}
		}
	}

	return sum
}

func (b *board) winner() bool {
	// horizontals
	for _, row := range b.spaces {
		winner := true
		for _, column := range row {
			if !column.filled {
				winner = false
				break
			}
		}

		if winner {
			return true
		}
	}

	// verticals
	for idx := 0; idx < 5; idx++ {
		winner := true
		for _, row := range b.spaces {
			if !row[idx].filled {
				winner = false
				break
			}
		}

		if winner {
			return true
		}
	}

	// diagonal1
	winner := false
	for x := 0; x < 5; x++ {
		row := b.spaces[x]
		column := row[x]

		if !column.filled {
			winner = false
			break
		}
	}

	if winner {
		return true
	}

	// diagonal2
	winner = false
	for x := 0; x < 5; x++ {
		row := b.spaces[x]
		column := row[4-x]

		if !column.filled {
			winner = false
			break
		}
	}

	if winner {
		return true
	}

	return false
}

type order []int

func NewOrder(line string) (order, error) {
	o := order{}

	values := strings.Split(line, ",")
	for _, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value: %w", err)
		}

		o = append(o, value)
	}

	return o, nil
}
