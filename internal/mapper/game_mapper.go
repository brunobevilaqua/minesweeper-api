package mapper

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
)

func MapModelToGameDto(g model.Game) (dto dto.GameDto) {
	dto.Id = g.Id
	dto.PlayerName = g.PlayerName
	dto.Status = string(g.Status)
	return
}
