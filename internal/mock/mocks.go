package mocks

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/model"
	"minesweeper-api/pkg/errors"
)

type SearchServiceMock struct {
	OnFindByGameId    func(id string) (*dto.GameResponse, *errors.ApiError)
	OnFindBoardById   func(id string) (*model.Board, *errors.ApiError)
	OnFindGame        func(id string) (*model.Game, *errors.ApiError)
	OnFindGameAndBord func(gameId string) (*model.Game, *model.Board, *errors.ApiError)
}

func (s *SearchServiceMock) FindByGameId(id string) (*dto.GameResponse, *errors.ApiError) {
	return s.OnFindByGameId(id)
}

func (s *SearchServiceMock) FindBoardById(id string) (*model.Board, *errors.ApiError) {
	return s.OnFindBoardById(id)
}

func (s *SearchServiceMock) FindGame(id string) (*model.Game, *errors.ApiError) {
	return s.OnFindGame(id)
}

func (s *SearchServiceMock) FindGameAndBord(gameId string) (*model.Game, *model.Board, *errors.ApiError) {
	return s.OnFindGameAndBord(gameId)
}

type MaintenanceServiceMock struct {
	OnAction        func(id string, d dto.ActionRequest) (*dto.GameResponse, *errors.ApiError)
	OnCreateNewGame func(d dto.CreateNewGameRequest) (*dto.GameResponse, *errors.ApiError)
}

func (m *MaintenanceServiceMock) PerformAction(id string, d dto.ActionRequest) (*dto.GameResponse, *errors.ApiError) {
	return m.OnAction(id, d)
}

func (m *MaintenanceServiceMock) CreateNewGame(d dto.CreateNewGameRequest) (*dto.GameResponse, *errors.ApiError) {
	return m.OnCreateNewGame(d)
}

type RepositoryMock struct {
	OnSaveGame        func(model.Game) error
	OnSaveBoard       func(model.Board) error
	OnFindGameById    func(id string) (*model.Game, error)
	OnFindBoardById   func(id string) (*model.Board, error)
	OnDeleteBoardById func(id string)
	OnDeleteGameById  func(id string)
}

func (r RepositoryMock) SaveGame(g model.Game) error {
	return r.OnSaveGame(g)
}

func (r RepositoryMock) SaveBoard(b model.Board) error {
	return r.OnSaveBoard(b)
}
func (r RepositoryMock) FindGameById(id string) (*model.Game, error) {
	return r.OnFindGameById(id)
}
func (r RepositoryMock) FindBoardById(id string) (*model.Board, error) {
	return r.OnFindBoardById(id)
}
func (r RepositoryMock) DeleteBoardById(id string) {
	r.OnDeleteBoardById(id)
}
func (r RepositoryMock) DeleteGameById(id string) {
	r.OnDeleteGameById(id)
}
