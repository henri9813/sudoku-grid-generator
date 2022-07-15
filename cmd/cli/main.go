package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sudoku_grid_generator/internal/grid"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Short: "hello",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}

		numberOfEmptyCells, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if numberOfEmptyCells <= 0 || numberOfEmptyCells >= 89 {
			return errors.New("the number of empty cells should be between 1 and 89")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		numberOfEmptyCells, err := strconv.Atoi(args[0])
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%s", err)
		}

		grid := grid.Generate(numberOfEmptyCells)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		for _, row := range grid {
			t.AppendRow(getPrettyfiedRow(row))
		}

		t.Render()
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getPrettyfiedRow(array []int) table.Row {
	str := table.Row{}
	for _, x := range array {
		value := strconv.Itoa(x)
		if x == 0 {
			value = ""
		}

		str = append(str, value)
	}

	return str
}
