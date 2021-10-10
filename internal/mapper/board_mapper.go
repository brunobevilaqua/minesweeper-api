package mapper

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
)

func MapModelToBoardDto(b model.Board) (dto dto.BoardDto) {
	dto.Columns = b.Cols
	dto.Rows = b.Rows
	dto.NumberOfMines = b.NumberOfMines
	return
}
