package mapper

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
	"strconv"
)

func MapModelToGameBoardDto(status model.GameStatus, b model.Board) dto.GameBoardDto {
	var showMines bool
	board := [][]string{}

	if status == model.GAME_STATUS_LOST {
		showMines = true
	}

	for r := 0; r < b.Rows; r++ {
		var row []string
		for c := 0; c < b.Cols; c++ {
			cell := b.Grid[r][c]

			if cell.Status == model.CELL_UNCLIKED {
				row = append(row, "_")
				continue
			}

			if cell.Status == model.CELL_CLICKED {
				if cell.NumberOfNearbyMines == 0 {
					row = append(row, "*")
					continue
				}
				row = append(row, strconv.Itoa(cell.NumberOfNearbyMines))
				continue
			}

			if cell.Status == model.CELL_FLAGGED {
				row = append(row, "F")
				continue
			}

			if showMines && cell.Mine {
				row = append(row, "M")
				continue
			}
		}
		board = append(board, row)
	}

	return dto.GameBoardDto{
		LabelDescription: dto.NewLabelDescritpio(),
		Board:            board,
		Status:           string(status),
	}
}
