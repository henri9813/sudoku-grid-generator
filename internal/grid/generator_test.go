package grid

import "testing"

func TestGenerate(t *testing.T) {
	grid := Generate(50)

	emptyCellsCount := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				emptyCellsCount++
			}
		}
	}

	if emptyCellsCount != 50 {
		t.Fatalf("unable to get 50 blank case, got: '%d'", emptyCellsCount)
	}
}
