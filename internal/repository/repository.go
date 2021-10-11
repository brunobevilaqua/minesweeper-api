package repository

import "minesweeper-api/internal/model"

type Repository interface {
	SaveGame(model.Game) error
	SaveBoard(model.Board) error
	FindGameById(id string) (*model.Game, error)
	FindBoardById(id string) (*model.Board, error)
	DeleteBoardById(id string)
	DeleteGameById(id string)
}
