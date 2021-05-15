package service

import (
	"project/servicelogs/pkg/controller"
	"project/servicelogs/pkg/repository"
)

type StatisticsService struct {
	repository repository.Statistics
}

func NewstatisticsService(repo repository.Statistics) *StatisticsService {
	return &StatisticsService{
		repository: repo,
	}
}

func (s *StatisticsService) GetStatistics() (controller.StatisticsResponse, error) {
	return s.repository.GetStatistics()
}
