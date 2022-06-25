package grid

import (
	"fmt"
	"golang.org/x/exp/slices"
)

//Grid represent the 9x9 grid
type Grid [][]int

//Coordinate represent the coordinate of a cell
type Coordinate [2]int

// To be simplified -> automatically computed
var squareCellsList = [][]Coordinate{
	{ //1
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	},
	{ //2
		{0, 3},
		{0, 4},
		{0, 5},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 4},
		{2, 5},
	},
	{ //3
		{0, 6},
		{0, 7},
		{0, 8},
		{1, 6},
		{1, 7},
		{1, 8},
		{2, 6},
		{2, 7},
		{2, 8},
	},

	{ //4
		{3, 0},
		{3, 1},
		{3, 2},
		{4, 0},
		{4, 1},
		{4, 2},
		{5, 0},
		{5, 1},
		{5, 2},
	},

	{
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 3},
		{4, 4},
		{4, 5},
		{5, 3},
		{5, 4},
		{5, 5},
	},
	{
		{3, 6},
		{3, 7},
		{3, 8},
		{4, 6},
		{4, 7},
		{4, 8},
		{5, 6},
		{5, 7},
		{5, 8},
	},
}

//IsValidGrid permit to say if the grid is valid or not.
func IsValidGrid(grid Grid) bool {
	if !grid.hasValidSize() {
		return false
	}

	if !grid.hasValidContent() {
		return false
	}

	if err := grid.verifyRows(); err != nil {
		return false
	}

	return true
}

func (grid Grid) hasValidSize() bool {
	if len(grid) != 9 {
		return false
	}

	for _, row := range grid {
		if len(row) != 9 {
			return false
		}
	}

	return true
}

func (grid Grid) hasValidContent() bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell < 1 || cell > 9 {
				return false
			}
		}
	}

	return true
}

func (grid Grid) verifyRows() error {
	for _, row := range grid {
		for j, cell := range row {
			if slices.Index(row, j+1) == -1 {
				return fmt.Errorf(
					"row %d has no '%d'",
					j+1,
					cell)
			}
		}
	}

	return nil
}

func (grid Grid) verifyCols() error {
	for column := 0; column < 9; column++ {
		founded := map[int]bool{}

		for _, row := range grid {
			if _, alreadyPresent := founded[row[column]]; alreadyPresent {
				return fmt.Errorf("duplicate %d in the array", row[column])
			}

			founded[row[column]] = true
		}
	}

	return nil
}

func (grid Grid) verifySquares() error {
	for _, squareCells := range squareCellsList {
		founded := map[int]bool{}

		for _, coordinate := range squareCells {
			if _, alreadyPresent := founded[grid[coordinate[0]][coordinate[1]]]; alreadyPresent {
				return fmt.Errorf("duplicate %d in the array", grid[coordinate[0]][coordinate[1]])
			}

			founded[grid[coordinate[0]][coordinate[1]]] = true
		}
	}

	return nil
}
