package grid

import "testing"

var goodGrid = Grid{
	0: {8, 1, 3, 9, 2, 5, 7, 4, 6},
	1: {9, 5, 6, 8, 4, 7, 3, 1, 2},
	2: {4, 7, 2, 3, 6, 1, 8, 9, 5},
	3: {6, 2, 4, 7, 1, 9, 5, 3, 8},
	4: {7, 9, 5, 6, 3, 8, 4, 2, 1},
	5: {3, 8, 1, 4, 5, 2, 9, 6, 7},
	6: {2, 3, 8, 1, 7, 4, 6, 5, 9},
	7: {5, 4, 9, 2, 8, 6, 1, 7, 3},
	8: {1, 6, 7, 5, 9, 3, 2, 8, 4},
}

var badGrid = map[string]Grid{
	"duplicatedNumberInCols": {
		0: {8, 1, 3, 9, 2, 5, 7, 4, 6},
		1: {8, 5, 6, 9, 4, 7, 3, 1, 2},
		2: {4, 7, 2, 3, 6, 1, 8, 9, 5},
		3: {6, 2, 4, 7, 1, 9, 5, 3, 8},
		4: {7, 9, 5, 6, 3, 8, 4, 2, 1},
		5: {3, 8, 1, 4, 5, 2, 9, 6, 7},
		6: {2, 3, 8, 1, 7, 4, 6, 5, 9},
		7: {5, 4, 9, 2, 8, 6, 1, 7, 3},
		8: {1, 6, 7, 5, 9, 3, 2, 8, 4},
	},
	"badSquareValues": { // Double 8 in the first square
		0: {8, 1, 3, 9, 2, 5, 7, 4, 6},
		1: {8, 5, 6, 9, 4, 7, 3, 1, 2},
		2: {4, 7, 2, 3, 6, 1, 8, 9, 5},
		3: {6, 2, 4, 7, 1, 9, 5, 3, 8},
		4: {7, 9, 5, 6, 3, 8, 4, 2, 1},
		5: {3, 8, 1, 4, 5, 2, 9, 6, 7},
		6: {2, 3, 8, 1, 7, 4, 6, 5, 9},
		7: {5, 4, 9, 2, 8, 6, 1, 7, 3},
		8: {1, 6, 7, 5, 9, 3, 2, 8, 4}},
}

func TestIsValidGridInvalidSize(t *testing.T) {
	var grid Grid
	grid = Grid{
		0: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
	}

	if IsValidGrid(grid) {
		t.Fatal("this grid should be not valid")
	}

	grid = Grid{
		0: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		1: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		2: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		3: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		4: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		5: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		6: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		7: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		8: {1, 2, 3, 4, 5, 6, 7, 8},
	}

	if IsValidGrid(grid) {
		t.Fatal("this grid should not be valid")
	}
}

func TestIsValidGridBadNumbers(t *testing.T) {
	grid := Grid{
		0: {
			10, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		1: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		2: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		3: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		4: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		5: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		6: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		7: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		8: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
	}

	if IsValidGrid(grid) {
		t.Fatal("a cell should be between 1 and 9 or be equal to -1")
	}
}

func TestIsValidGridDuplicateNumberInARow(t *testing.T) {
	grid := Grid{
		0: {
			2, 2, 2, 4, 5, 6, 7, 8, 9,
		},
		1: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		2: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		3: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		4: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		5: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		6: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		7: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		8: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
	}

	if IsValidGrid(grid) {
		t.Fatal("a cell should be between 1 and 9 or be equal to -1")
	}
}

func TestIsValidGrid(t *testing.T) {
	grid := Grid{
		0: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		1: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		2: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		3: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		4: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		5: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		6: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		7: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
		8: {
			1, 2, 3, 4, 5, 6, 7, 8, 9,
		},
	}

	if !IsValidGrid(grid) {
		t.Fatal("the grid should be valid")
	}
}

func TestGridVerifyRows(t *testing.T) {
	if err := goodGrid.verifyCols(); err != nil {
		t.Fatalf("this grid should be valid, got: '%v'", err)
	}

	if err := badGrid["duplicatedNumberInCols"].verifyCols(); err == nil {
		t.Fatal("a grid with equal number in columns should be invalid")
	}
}

func TestGridVerifySquares(t *testing.T) {
	if err := goodGrid.verifySquares(); err != nil {
		t.Fatalf("this grid should be valid, got: '%v'", err)
	}

	if err := badGrid["badSquareValues"].verifySquares(); err == nil {
		t.Fatal("a grid with equal number in columns should be invalid")
	}
}
