package dto

import (
	"minesweeper-api/internal/model"
	"time"
)

type GameDto struct {
	Id         string    `json:"id"`
	PlayerName string    `json:"playerName"`
	Status     string    `json:"status"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

func NewGameDto(gameId string, playerName string, status model.GameStatus, startTime, endTime time.Time) GameDto {
	dto := GameDto{}
	dto.Status = status.GetString()
	dto.Id = gameId
	dto.PlayerName = playerName
	dto.StartTime = startTime
	dto.EndTime = endTime
	return dto
}
