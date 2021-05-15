package service

import (
	"project/servicelogs/pkg/controller"
	"project/servicelogs/pkg/repository"
)

type Statistics interface {
	GetStatistics() (controller.StatisticsResponse, error)
}

type Service struct {
	Statistics
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Statistics: NewstatisticsService(r.Statistics),
	}
}
