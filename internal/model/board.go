package model

import (
	"math/rand"
)

// Game Board Constants
const (
	MAX_DEFAULT_ROWS  = 30
	MAX_DEFAULT_COLS  = 30
	MIN_DEFAULT_ROWS  = 5
	MIN_DEFAULT_COLS  = 5
	MIN_DEFAULT_MINES = 8
	MAX_DEFAULT_MINES = 32
)

type CellStatus string

// Cell Status Constants
const (
	CELL_UNCLIKED CellStatus = "Not Clicked"
	CELL_CLICKED  CellStatus = "Clicked"
	CELL_FLAGGED  CellStatus = "Flagged"
	CELL_EXPLODED CellStatus = "Exploded"
	CELL_EXPANDED CellStatus = "Expanded"
)

type Position struct {
	Row int `json:"x"`
	Col int `json:"y"`
}

type Board struct {
	Rows            int        `json:"rows"`
	Cols            int        `json:"columns"`
	NumberOfMines   int        `json:"mines"`
	MinesDiscovered int        `json:"minesDiscovered"`
	MinesPositions  []Position `json:"minesCoordinates"`
	Grid            []CellGrid `json:"grid"`
	GameId          string     `json:"gameId"`
	Clicks          int        `json:"clicks"`
	Flags           int        `json:"flags"`
}

type CellGrid []Cell

type Cell struct {
	NearbyCells         []Position `json:"nearbyCells"`
	NumberOfNearbyMines int        `json:"numberOfNearbyMines"`
	Status              CellStatus `json:"status"`
	Pos                 Position   `json:"position"`
	Mine                bool       `json:"mine"`
	Evaluated           bool       `json:"evaluated"`
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

func NewBoard(rows, cols, mines int, gameId string) Board {
	totalOfRows := setValue(rows, MAX_DEFAULT_ROWS, MIN_DEFAULT_ROWS)
	totalOfColumns := setValue(cols, MAX_DEFAULT_COLS, MIN_DEFAULT_COLS)
	totalOfMines := setValue(cols, MAX_DEFAULT_MINES, MIN_DEFAULT_MINES)

	b := Board{Cols: totalOfColumns, Rows: totalOfRows, NumberOfMines: totalOfMines, GameId: gameId}

	// Building Board
	for r := 0; r < totalOfRows; r++ {
		var row []Cell
		for c := 0; c < totalOfColumns; c++ {
			cell := Cell{Pos: Position{Row: r, Col: c}, Status: CELL_UNCLIKED}
			row = append(row, cell)
		}
		b.Grid = append(b.Grid, row)
	}

	// Settings Randomly adding mines...
	m := 0
	for m < totalOfMines {
		row := rand.Intn(totalOfRows - 1)
		col := rand.Intn(totalOfColumns - 1)

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
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c])
					continue
				} else if c == totalOfColumns-1 { // first row + last column
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c])
					continue
				} else { // first row only
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c+1])
					continue
				}
			}

			// last row
			if r == totalOfRows-1 {
				if c == 0 { // last row + first column
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c+1])
					continue
				} else if c == totalOfColumns-1 { // last row + last column
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c-1])
					continue
				} else { // last row
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c+1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c-1])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c])
					currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c+1])
					continue
				}
			}

			// first column
			if c == 0 {
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c+1])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c+1])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c+1])
				continue
			}

			// last column
			if c == totalOfColumns-1 {
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c-1])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c-1])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c])
				currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c-1])
				continue
			}

			// remaining cells... midle ones...
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c-1])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c-1])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c-1])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r-1][c+1])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r][c+1])
			currentCell.setNearbyAndUpdateMinesCountrer(b.Grid[r+1][c+1])
		}
	}
	return b
}

func (c *Cell) setNearbyAndUpdateMinesCountrer(cell Cell) {
	if cell.Mine {
		c.NumberOfNearbyMines += 1
	}
	c.NearbyCells = append(c.NearbyCells, cell.Pos)
}

func (b *Board) ExpandNearbyCell(currentCell Cell) {
	if !currentCell.Evaluated {
		// cell := &b.Grid[currentCell.Pos.Row][currentCell.Pos.Col]
		b.Grid[currentCell.Pos.Row][currentCell.Pos.Col].Evaluated = true
		// cell.Evaluated = true
		// if !cell.Mine && cell.NumberOfNearbyMines == 0 {
		if !b.Grid[currentCell.Pos.Row][currentCell.Pos.Col].Mine && b.Grid[currentCell.Pos.Row][currentCell.Pos.Col].NumberOfNearbyMines == 0 {
			b.Grid[currentCell.Pos.Row][currentCell.Pos.Col].Status = CELL_EXPANDED

			for i := 0; i < len(currentCell.NearbyCells); i++ {
				x := currentCell.NearbyCells[i].Row
				y := currentCell.NearbyCells[i].Col

				if b.Grid[x][y].Evaluated {
					continue
				}
				b.Grid[x][y].Evaluated = true
				if !b.Grid[x][y].Mine && b.Grid[x][y].NumberOfNearbyMines == 0 {
					b.Grid[x][y].Status = CELL_EXPANDED
					continue
				}
			}
			// for _, position := range currentCell.NearbyCells {
			// 	nearbyCell := &b.Grid[position.Row][position.Col]
			// 	if nearbyCell.Evaluated {
			// 		continue
			// 	}
			// 	nearbyCell.Evaluated = true
			// 	if !nearbyCell.Mine && nearbyCell.NumberOfNearbyMines == 0 {
			// 		nearbyCell.Status = CELL_EXPANDED
			// 		continue
			// 	}
			// }
		}
	}
}

func (b *Board) BoardEnded() bool {
	return b.MinesDiscovered == b.NumberOfMines
}
