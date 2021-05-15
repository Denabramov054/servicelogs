package repository

import (
	"project/servicelogs/pkg/controller"

	"github.com/jmoiron/sqlx"
)

type Statistics interface {
	GetStatistics() (controller.StatisticsResponse, error)
}

type Repository struct {
	Statistics
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Statistics: NewStatistics(db),
	}
}
