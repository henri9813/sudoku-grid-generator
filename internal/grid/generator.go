package grid

import (
	"math/rand"
	"sudoku_grid_generator/pkg/utils"
	"time"
)

func Generate(emptyCellsCount int) Grid {
retry:
	var grid = Grid{}
	_ = grid.initializeEmpty()

	//Generate a valid grid with good cells
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			possibilities := getPossibleCellValues(grid, i, j)

			if len(possibilities) == 0 {
				goto retry
			}

			rand.Seed(time.Now().UnixNano())
			grid[i][j] = possibilities[rand.Intn(len(possibilities))]
		}
	}

	// Remove random number and verify the grid is still solvable !!
	for x := 0; x < emptyCellsCount; {
		coordinate := Coordinate{
			utils.RandInt(0, 8),
			utils.RandInt(0, 8),
		}

		// Already a removed cell ;-)
		if grid[coordinate[0]][coordinate[1]] == 0 {
			continue
		}

		grid[coordinate[0]][coordinate[1]] = 0
		x++
	}

	return grid
}

func getPossibleCellValues(grid Grid, i int, j int) (possibilities []int) {
	possibilities = []int{}
	tmpGrid := grid
	for number := 1; number < 10; number++ {
		tmpGrid[i][j] = number

		if err := tmpGrid.verifyRows(); err != nil {
			continue
		}

		if err := tmpGrid.verifyCols(); err != nil {
			continue
		}

		if err := tmpGrid.verifySquares(); err != nil {
			continue
		}

		possibilities = append(possibilities, number)
	}

	return possibilities
}
