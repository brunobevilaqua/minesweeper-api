package service

import "minesweeper-api/internal/repository"

type Service struct {
	Search      SearchService
	Maintenance MaintenanceService
}

func NewService(repository repository.Repository) Service {
	s := Service{Search: NewSearchService(repository),
		Maintenance: NewMaintenanceService(repository)}

	return s
}
