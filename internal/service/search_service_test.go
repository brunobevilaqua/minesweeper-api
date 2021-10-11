package service_test

import (
	"fmt"
	mocks "minesweeper-api/internal/mock"
	"minesweeper-api/internal/model"
	"minesweeper-api/internal/service"
	"testing"
)

func TestFindByGameId(t *testing.T) {
	svc := service.SearchService{Store: &mocks.RepositoryMock{
		OnFindGameById: func(id string) (*model.Game, error) {
			return &model.Game{PlayerName: "Bruno", Id: "game", Status: model.GAME_STATUS_PLAYING}, nil
		},
		OnFindBoardById: func(id string) (*model.Board, error) {
			return &model.Board{}, nil
		},
	},
	}

	response, _ := svc.FindByGameId("123")

	if response.Data.GameDto.Id != "game" {
		t.Errorf("unexpected id")
	}
}

func TestFindByGameId_NotFound(t *testing.T) {
	svc := service.SearchService{Store: &mocks.RepositoryMock{
		OnFindGameById: func(id string) (*model.Game, error) {
			return nil, fmt.Errorf("%T", 404)
		},
		OnFindBoardById: func(id string) (*model.Board, error) {
			return nil, fmt.Errorf("%T", 404)
		},
	},
	}

	_, err := svc.FindByGameId("123")

	if err == nil {
		t.Errorf("should return error")
	}
}

func TestFindBoardByGameId(t *testing.T) {
	svc := service.SearchService{Store: &mocks.RepositoryMock{
		OnFindBoardById: func(id string) (*model.Board, error) {
			return &model.Board{GameId: "123"}, nil
		},
	},
	}

	response, _ := svc.FindBoardById("123")

	if response.GameId != "123" {
		t.Errorf("unexpected id")
	}
}

func TestFindBoardByGameId_NotFound(t *testing.T) {
	svc := service.SearchService{Store: &mocks.RepositoryMock{
		OnFindBoardById: func(id string) (*model.Board, error) {
			return nil, fmt.Errorf("%T", 404)
		},
	},
	}

	_, err := svc.FindBoardById("123")

	if err == nil {
		t.Errorf("should return error")
	}
}

func TestFindByGameAndBoard(t *testing.T) {
	svc := service.SearchService{Store: &mocks.RepositoryMock{
		OnFindGameById: func(id string) (*model.Game, error) {
			return &model.Game{PlayerName: "Bruno", Id: "123", Status: model.GAME_STATUS_PLAYING}, nil
		},
		OnFindBoardById: func(id string) (*model.Board, error) {
			return &model.Board{GameId: "123"}, nil
		},
	},
	}

	game, board, _ := svc.FindGameAndBord("123")

	if game.Id != "123" && board.GameId != "123" {
		t.Errorf("Should return game and board models")
	}
}

func TestFindByGameAndBoard_NotFound(t *testing.T) {
	svc := service.SearchService{Store: &mocks.RepositoryMock{
		OnFindGameById: func(id string) (*model.Game, error) {
			return nil, fmt.Errorf("%T", 404)
		},
		OnFindBoardById: func(id string) (*model.Board, error) {
			return nil, fmt.Errorf("%T", 404)
		},
	},
	}

	game, board, err := svc.FindGameAndBord("123")

	if game != nil && board != nil {
		t.Errorf("should return error")
	}

	if err == nil {
		t.Errorf("should return error")
	}
}
