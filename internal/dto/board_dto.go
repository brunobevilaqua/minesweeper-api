package dto

import "minesweeper-api/internal/model"

type BoardDto struct {
	Columns          int              `json:"columns"`
	Rows             int              `json:"rows"`
	NumberOfMines    int              `json:"numberOfMines"`
	Clicks           int              `json:"clicks"`
	DiscoveredMines  int              `json:"discoveredMines"`
	Cells            []CellDto        `json:"grid"`
	MinesCoordinates []model.Position `json:"minesCoordinates"`
}

type CellDto struct {
	Status              model.CellStatus `json:"status"`
	Position            model.Position   `json:"position"`
	NumberOfNearbyMines int              `json:"numberOfNearbyMines"`
	Mine                bool             `json:"mine"`
}

func NewBoardDto(row, col, numberOfMines, clicks, discoveredMines int, grid []model.CellGrid, minesCoordinates []model.Position) BoardDto {
	cells := []CellDto{}

	for _, row := range grid {
		for _, cell := range row {
			cells = append(cells, NewCellDto(cell.Status, cell.Pos, cell.NumberOfNearbyMines, cell.Mine))
		}
	}

	return BoardDto{
		Columns:          col,
		Rows:             row,
		NumberOfMines:    numberOfMines,
		Clicks:           clicks,
		DiscoveredMines:  discoveredMines,
		Cells:            cells,
		MinesCoordinates: minesCoordinates,
	}
}

func NewCellDto(status model.CellStatus, position model.Position, numberOfNearbyMines int, mine bool) CellDto {
	return CellDto{
		Status:              status,
		Position:            position,
		NumberOfNearbyMines: numberOfNearbyMines,
		Mine:                mine,
	}
}
