package model

import (
	"minesweeper-api/pkg/errors"
	"time"
)

type GameStatus int

const (
	GAME_STATUS_PLAYING GameStatus = iota
	GAME_STATUS_LOST
	GAME_STATUS_WON
)

type Game struct {
	PlayerName string     `json:"PlayerName"`
	StartTime  time.Time  `json:"startTime"`
	EndTime    time.Time  `json:"endTime"`
	Status     GameStatus `json:"status"`
	Id         string     `json:"id"`
}

func NewGame(playerName string, rows, cols, mines int, gameId string) (*Game, *errors.ApiError) {
	if playerName == "" {
		return nil, errors.NewApiError(errors.INVALID_USER_NAME_ERROR)
	}

	return &Game{
		PlayerName: playerName,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Id:         gameId,
		Status:     GAME_STATUS_PLAYING,
	}, nil
}

func (s GameStatus) GetString() string {
	switch s {
	case GAME_STATUS_LOST:
		return "Game Over - You Lost!"
	case GAME_STATUS_WON:
		return "You Won - Congratulations!"
	default:
		return "Still Playing..."
	}
}
