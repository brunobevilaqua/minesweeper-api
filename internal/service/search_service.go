package service

import (
	"minesweeper-api/internal/dto"
	"minesweeper-api/internal/repository"
	"minesweeper-api/pkg/errors"
)

type SearchServiceInterface interface {
	FindByPlayerName(name string) (*dto.Response, *errors.ApiError)
	FindByGameId(gameId string) (*dto.Response, *errors.ApiError)
}

type SearchService struct {
	store repository.Repository
}

func NewSearchService(r repository.Repository) SearchService {
	return SearchService{store: r}
}

func (s SearchService) FindByGameId(id string) (*dto.Response, *errors.ApiError) {
	if id == "" {
		return nil, errors.NewApiError(errors.InvalidParameter)
	}

	game, err := s.store.FindById(id)

	if err != nil {
		return nil, errors.NewApiError(errors.NoRecordsFound)
	}

	return dto.NewResponse(*game), nil
}

func (s SearchService) FindByPlayerName(name string) (*dto.Response, *errors.ApiError) {
	if name == "" {
		return nil, errors.NewApiError(errors.InvalidUserName)
	}

	game, err := s.store.FindByPlayerName(name)

	if err != nil {
		return nil, errors.NewApiError(errors.NoRecordsFound)
	}

	return dto.NewResponse(*game), nil
}
