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
	api := router.Group("/api")

	api.GET("/games/:id", controller.Search.FindByGameId)

	api.POST("/games", func(c *gin.Context) {
		controller.Maintenance.CreateNewGame(*c)
	})
	api.PUT("/games/:id", func(c *gin.Context) {
		controller.Maintenance.Click(*c)
	})

	router.Run()

	defer redis.Close()
}
