package dto

import (
	"minesweeper-api/internal/model"
	"time"
)

type GameResponse struct {
	Data struct {
		GameDto  GameDto  `json:"game"`
		BoardDto BoardDto `json:"board"`
	} `json:"data"`
}

type GameUpdateResponse struct {
	Data struct {
		GameBoardDto `json:"board"`
	}
}

type GameResultsResponse struct {
	Data struct {
		GameStatus   string `json:"gameStatus"`
		GameBoardDto `json:"board"`
		StartTime    time.Time `json:"startTime"`
		EndTime      time.Time `json:"endTime"`
		PlayerName   string    `json:"playerName"`
		Clicks       int       `json:"clicks"`
	}
}

func NewGameResultsResponse(status model.GameStatus, startTime, endTime time.Time, playerName string, clicks int, dto GameBoardDto) GameResultsResponse {
	response := GameResultsResponse{}

	response.Data.GameStatus = status.GetString()
	response.Data.StartTime = startTime
	response.Data.EndTime = endTime
	response.Data.PlayerName = playerName
	response.Data.Clicks = clicks
	response.Data.GameBoardDto = dto
	return response
}

func NewGameUpdateResponse(dto GameBoardDto) GameUpdateResponse {
	response := GameUpdateResponse{}
	response.Data.GameBoardDto = dto
	return response
}

func NewGameResponse(boardDto BoardDto, gameDto GameDto) GameResponse {
	response := GameResponse{}
	response.Data.BoardDto = boardDto
	response.Data.GameDto = gameDto
	return response
}
