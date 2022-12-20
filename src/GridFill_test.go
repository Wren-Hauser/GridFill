package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitBoilerPlate(t *testing.T) {

	tests := []struct {
		name     string
		in       Input
		expected int
		validate func(tc *testing.T, expected, actual int) bool
	}{
		{
			"test the tests",
			Input{
				1,
				1,
				[]Pos{{1, 1}, {2, 2}},
			},
			1,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.validate(t, test.expected, 1)
		})
	}
}

func TestUnitGetSize(t *testing.T) {

	tests := []struct {
		name     string
		in       Input
		expected int
		validate func(tc *testing.T, expected, actual int) bool
	}{
		{
			"base case- get size of square grid, no holes",
			Input{
				3,
				3,
				[]Pos{},
			},
			9,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base case- get size of square grid, two holes",
			Input{
				3,
				3,
				[]Pos{{1, 1}, {2, 2}},
			},
			7,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.validate(t, test.expected, getSize(test.in))
		})
	}
}

func TestUnitGetRuns(t *testing.T) {

	tests := []struct {
		name     string
		in       Input
		n        int
		expected int
		validate func(tc *testing.T, expected, actual int) bool
	}{
		{
			"base case- square board of n size, no holes",
			Input{
				4,
				4,
				[]Pos{},
			},
			4,
			8,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base case- square board of n size, 1 holes",
			Input{
				4,
				4,
				[]Pos{{0, 0}},
			},
			4,
			6,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base case- square board of n size, all holes",
			Input{
				2,
				2,
				[]Pos{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			},
			2,
			0,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base case- rectangular board of n, n - 1 size, no holes",
			Input{
				4,
				3,
				[]Pos{},
			},
			4,
			3,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base case- square board of n-1 size, no holes",
			Input{
				3,
				3,
				[]Pos{},
			},
			4,
			0,
			func(tc *testing.T, expected, actual int) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			grid := createGrid(test.in)

			test.validate(t, test.expected, len(getRuns(grid, test.n)))
		})
	}
}

func TestUnitGridFill(t *testing.T) {

	tests := []struct {
		name     string
		in       Input
		expected []Run
		validate func(tc *testing.T, expected, actual []Run) bool
	}{
		{
			"base- recursive base case 0x0 grid, no holes",
			Input{
				0,
				0,
				[]Pos{},
			},
			nil,
			func(tc *testing.T, expected, actual []Run) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base- 3x3 grid, no holes",
			Input{
				3,
				3,
				[]Pos{},
			},
			[]Run{
				{Pos{0, 0}, Pos{0, 2}},
				{Pos{1, 0}, Pos{1, 2}},
				{Pos{2, 0}, Pos{2, 2}}},
			func(tc *testing.T, expected, actual []Run) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base- 6x3 grid, no holes",
			Input{
				6,
				3,
				[]Pos{},
			},
			[]Run{
				{Pos{0, 0}, Pos{0, 2}},
				{Pos{1, 0}, Pos{1, 2}},
				{Pos{2, 0}, Pos{2, 2}},
				{Pos{3, 0}, Pos{3, 2}},
				{Pos{4, 0}, Pos{4, 2}},
				{Pos{5, 0}, Pos{5, 2}}},
			func(tc *testing.T, expected, actual []Run) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base- 6x4 grid, no holes, vertical and horizontal filling",
			Input{
				6,
				4,
				[]Pos{},
			},
			[]Run{
				{Pos{0, 0}, Pos{0, 2}},
				{Pos{1, 0}, Pos{1, 2}},
				{Pos{2, 0}, Pos{2, 2}},
				{Pos{3, 0}, Pos{3, 2}},
				{Pos{4, 0}, Pos{4, 2}},
				{Pos{5, 0}, Pos{5, 2}},
				{Pos{0, 3}, Pos{2, 3}},
				{Pos{3, 3}, Pos{5, 3}}},
			func(tc *testing.T, expected, actual []Run) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
		{
			"base- 3x4 grid, 3 holes",
			Input{
				3,
				4,
				[]Pos{
					{1, 1},
					{1, 2},
					{1, 3},
				},
			},
			[]Run{
				{Pos{0, 1}, Pos{0, 3}},
				{Pos{2, 1}, Pos{2, 3}},
				{Pos{0, 0}, Pos{2, 0}}},
			func(tc *testing.T, expected, actual []Run) bool {
				return assert.Equal(tc, expected, actual)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid := createGrid(test.in)
			var runs = getRuns(grid, 3)
			output := coverGrid(getSize(test.in), runs, nil)
			test.validate(t, test.expected, output)
		})
	}
}
