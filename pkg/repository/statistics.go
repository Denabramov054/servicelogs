package repository

import (
	"fmt"
	"project/servicelogs/pkg/controller"

	"github.com/jmoiron/sqlx"
)

type Statistic struct {
	db *sqlx.DB
}

func NewStatistics(db *sqlx.DB) *Statistic {
	return &Statistic{
		db: db,
	}
}

func (s *Statistic) GetStatistics() (controller.StatisticsResponse, error) {
	var inputs controller.StatisticsResponse
	query := fmt.Sprintf(`SELECT 
		user.id, user.name, count(likes.id) as count_likes, count(comments.id) as count_comments
		FROM %s user
		LEFT JOIN %s likes ON likes.user_id = user.id
		LEFT JOIN %s comments ON comments.user_id = user.id
		ORDER BY user.id`, userTable, likeTable, commentTable)

	if err := s.db.Select(&inputs, query); err != nil {
		return inputs, err
	}
	return inputs, nil
}
