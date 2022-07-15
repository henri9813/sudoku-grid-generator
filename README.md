# sudoku-grid-generator

Generator of sudoku grid

## Usage

```bash
go run cmd/cli/main.go <NUMBER_EMPTY_CELLS>
```

Example:

```bash
hdevigne@MacBook-Pro-de-Henri:~/GolandProjects/sudoku-grid-generator » go run cmd/cli/main.go 9
+---+---+---+---+---+---+---+---+---+
|   | 1 |   | 9 |   | 8 |   | 3 | 4 |
|   | 3 | 8 | 5 | 1 | 2 | 9 | 7 | 6 |
| 9 | 2 |   | 7 | 4 | 3 | 8 | 5 | 1 |
|   | 8 |   | 6 | 2 | 7 | 1 | 4 | 5 |
| 7 | 4 | 5 | 3 | 8 | 1 | 6 | 9 | 2 |
| 1 | 6 | 2 | 4 | 9 | 5 | 7 | 8 | 3 |
| 2 | 5 |   | 1 | 7 | 9 | 4 | 6 | 8 |
| 8 | 9 | 4 | 2 | 3 | 6 | 5 | 1 | 7 |
| 6 | 7 | 1 | 8 | 5 | 4 | 3 | 2 | 9 |
+---+---+---+---+---+---+---+---+---+
```
