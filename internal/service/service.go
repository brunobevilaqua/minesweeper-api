package service

import "minesweeper-api/internal/repository"

type Service struct {
	Search      SearchServiceInterface
	Maintenance MaintenanceServiceInterface
}

func NewService(repository repository.Repository) Service {
	searchService := NewSearchService(repository)
	maintenanceService := NewMaintenanceService(repository, searchService)

	s := Service{Search: searchService,
		Maintenance: maintenanceService}

	return s
}
