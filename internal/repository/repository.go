package repository

import (
	"minesweeper-api/internal/model"
)

type Repository interface {
	Save(g model.Game) (*model.Game, error)
	FindById(id string) (*model.Game, error)
	FindByPlayerName(name string) (*model.Game, error)
}
