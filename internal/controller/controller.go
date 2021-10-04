package controller

import "minesweeper-api/internal/service"

type Controller struct {
	Search      SearchController
	Maintenance MaintenanceController
}

func NewController(service service.Service) Controller {
	return Controller{Search: NewSearchController(service.Search),
		Maintenance: NewMaintenanceController(service.Maintenance)}
}
