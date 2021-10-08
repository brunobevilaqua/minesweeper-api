package model

import "math/rand"

// Game Board Constants
const (
	MAX_DEFAULT_ROWS  = 30
	MAX_DEFAULT_COLS  = 30
	MIN_DEFAULT_ROWS  = 8
	MIN_DEFAULT_COLS  = 8
	MIN_DEFAULT_MINES = 12
	MAX_DEFAULT_MINES = 12
)

// Cell Status Constants
const (
	CELL_UNCLIKED int = iota
	CELL_CLICKED
	CELL_FLAGGED
	CELL_EXPLODED
)

type Position struct {
	Row int `json:"x"`
	Col int `json:"y"`
}

type Board struct {
	Rows           int        `json:"rows"`
	Cols           int        `json:"columns"`
	NumberOfMines  int        `json:"mines"`
	MinesPositions []Position `json:"minesCoordinates"`
	Grid           []CellGrid `json:"grid"`
}

type CellGrid []Cell

type Cell struct {
	Evaluated           bool     `json:"evaluated"`
	NearbyCells         []*Cell  `json:"nearbyCells"`
	NumberOfNearbyMines int      `json:"numberOfNearbyMines"`
	Status              int      `json:"status"`
	Pos                 Position `json:"position"`
	Mine                bool     `json:"mine"`
}

func setValue(v, max, min int) int {
	// if input value is greater than MAX_DEFAULT
	if v > max {
		return max
	}

	// if input value is less than MIN_DEFAULT
	if v < min {
		return min
	}

	return v
}

func NewBoard(rows, cols, mines int) *Board {
	totalOfRows := setValue(rows, MAX_DEFAULT_ROWS, MIN_DEFAULT_ROWS)
	totalOfColumns := setValue(cols, MAX_DEFAULT_COLS, MIN_DEFAULT_COLS)
	totalOfMines := setValue(cols, MAX_DEFAULT_MINES, MIN_DEFAULT_MINES)

	b := Board{Cols: totalOfColumns, Rows: totalOfRows, NumberOfMines: totalOfMines}

	// Building Board
	for r := 0; r < totalOfRows; r++ {
		var row []Cell
		for c := 0; c < cols; c++ {
			cell := Cell{Pos: Position{Row: totalOfRows, Col: totalOfColumns}}
			row = append(row, cell)
		}
		b.Grid = append(b.Grid, row)
	}

	// Settings Randomly adding mines...
	m := 0
	for m < totalOfMines {
		row := rand.Intn(totalOfRows)
		col := rand.Intn(totalOfColumns)

		if !b.Grid[row][col].Mine {
			b.Grid[row][col].Mine = true
			b.MinesPositions = append(b.MinesPositions, Position{Row: row, Col: col})
			m++
		}
	}

	// Setting NearbyCells...
	for r := 0; r < totalOfRows; r++ {
		for c := 0; c < totalOfColumns; c++ {
			currentCell := &b.Grid[r][c]

			if currentCell.Mine {
				continue
			}

			// first row
			if r == 0 {
				if c == 0 { // first row + first column
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c])
					continue
				} else if c == totalOfColumns-1 { // first row + last column
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c])
					continue
				} else { // first row only
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c+1])
					continue
				}
			}

			// last row
			if r == totalOfRows-1 {
				if c == 0 { // last row + first column
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c+1])
					continue
				} else if c == totalOfColumns-1 { // last row + last column
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c-1])
					continue
				} else { // last row
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c+1])
					continue
				}
			}

			// first column
			if c == 0 {
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c+1])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c+1])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c+1])
				continue
			}

			// last column
			if c == totalOfColumns-1 {
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c-1])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c-1])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c-1])
				continue
			}

			// remaining cells... midle ones...
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c-1])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c-1])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c-1])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r-1][c+1])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r][c+1])
			currentCell.setNearbyAndUpdateMinesCountrer(&b.Grid[r+1][c+1])
		}
	}
	return &b
}

func (c *Cell) setNearbyAndUpdateMinesCountrer(cell *Cell) {
	if cell.Mine {
		c.NumberOfNearbyMines += 1
	}
	c.NearbyCells = append(c.NearbyCells, cell)
}
