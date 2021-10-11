package controller_test

import (
	"minesweeper-api/internal/controller"
	"minesweeper-api/internal/dto"
	mocks "minesweeper-api/internal/mock"
	"minesweeper-api/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFindGameById(t *testing.T) {
	controller := controller.SearchController{
		Service: &mocks.SearchServiceMock{
			OnFindByGameId: func(id string) (*dto.GameResponse, *errors.ApiError) {
				gameDto := dto.GameDto{
					Id: "123",
				}
				response := dto.GameResponse{}
				response.Data.GameDto = gameDto
				return &dto.GameResponse{}, nil
			},
		},
	}

	r := gin.Default()
	r.GET("/games/:id", controller.FindByGameId)
	r.Run(":8080")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/games/:id", nil)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Shoud return 200 when finding game")
	}
}

func TestFindGameById_InvalidId(t *testing.T) {
	controller := controller.SearchController{
		Service: &mocks.SearchServiceMock{
			OnFindByGameId: func(id string) (*dto.GameResponse, *errors.ApiError) {
				return nil, errors.NewApiError(errors.NO_RECORDS_FOUND_ERROR)
			},
		},
	}

	r := gin.Default()
	r.GET("/games/:id", controller.FindByGameId)
	r.Run(":8080")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/games/123", nil)
	r.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Fatalf("Shoud return 400 when game not found")
	}
}
