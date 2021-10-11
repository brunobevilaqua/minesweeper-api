package dto

import (
	"minesweeper-api/internal/model"
	"strconv"
)

type Label string

const (
	EMPTY_CELL   Label = "Empty Cell"
	NEARBY_MINES Label = "Number of nearby mines"
	MINE         Label = "Mine"
	FLAG         Label = "Flag"
	UNCLICKED    Label = "Unclicked Cell"
)

type LabelDescription struct {
	EmptyCell   Label `json:"X"`
	NearbyMines Label `json:"<Number>"`
	Mine        Label `json:"M"`
	Flag        Label `json:"F"`
	Unclicked   Label `json:"_"`
}

type GameBoardDto struct {
	LabelDescription `json:"labels"`
	Status           string   `json:"status"`
	Board            []string `json:"board"`
}

func NewLabelDescritpion() LabelDescription {
	return LabelDescription{
		EmptyCell:   EMPTY_CELL,
		NearbyMines: NEARBY_MINES,
		Mine:        MINE,
		Flag:        FLAG,
		Unclicked:   UNCLICKED,
	}
}

func NewGameBoardDto(status model.GameStatus, b model.Board) GameBoardDto {
	var showMines bool
	var board []string

	if status == model.GAME_STATUS_LOST {
		showMines = true
	}

	for r := 0; r < b.Rows; r++ {
		var row string
		for c := 0; c < b.Cols; c++ {
			cell := b.Grid[r][c]

			if c != 0 {
				row += "  "
			}

			if cell.Status == model.CELL_UNCLIKED {
				row += "_"
				continue
			}

			if cell.Status == model.CELL_CLICKED || cell.Status == model.CELL_EXPANDED {
				if cell.NumberOfNearbyMines == 0 {
					row += "X"
					continue
				}
				row += strconv.Itoa(cell.NumberOfNearbyMines)
				continue
			}

			if cell.Status == model.CELL_FLAGGED {
				row += "F"
				continue
			}

			if showMines && cell.Mine {
				row += "M"
				continue
			}
		}

		board = append(board, row)
	}

	return GameBoardDto{
		LabelDescription: NewLabelDescritpion(),
		Board:            board,
		Status:           string(status),
	}
}
