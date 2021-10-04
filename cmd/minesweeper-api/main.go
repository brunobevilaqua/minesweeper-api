package main

import (
	"minesweeper-api/internal/controller"
	"minesweeper-api/internal/repository"
	"minesweeper-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	redis := repository.NewRedisStore()
	service := service.NewService(redis)
	controller := controller.NewController(service)

	router := gin.Default()
	router.GET("/game:id", controller.Search.FindByGameId)
	router.GET("/game:name", controller.Search.FindByUserName)
	router.POST("/game", func(c *gin.Context) {
		controller.Maintenance.CreateNewGame(*c)
	})
	router.PUT("/game", func(c *gin.Context) {
		controller.Maintenance.Click(*c)
	})

	router.Run("localhost:8080")

	defer redis.Close()
}
