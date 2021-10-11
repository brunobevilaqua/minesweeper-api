package dto

import "minesweeper-api/internal/model"

type GameDto struct {
	Id         string `json:"id"`
	PlayerName string `json:"playerName"`
	Status     string `json:"status"`
}

func NewGameDto(gameId string, playerName string, status model.GameStatus) GameDto {
	dto := GameDto{}
	dto.Status = status.GetString()
	dto.Id = gameId
	dto.PlayerName = playerName
	return dto
}
