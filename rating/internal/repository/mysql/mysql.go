package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-microservice/rating/internal/repository"
	"go-microservice/rating/pkg/model"
)

type Repository struct {
	db *sql.DB
}

func New() (*Repository, error) {
	db, err := sql.Open("mysql", "root:1234@/movieexample")
	if err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}

func (r *Repository) Get(ctx context.Context, recordId model.RecordId, recordType model.RecordType) ([]model.Rating, error) {
	query := `SELECT user_id,value FROM ratings WHERE record_id = ? AND record_type = ?`
	rows, err := r.db.QueryContext(ctx, query, recordId, recordType)
	if err != nil {
		return nil, err

	}

	defer rows.Close()
	var res []model.Rating

	for rows.Next() {
		var userID string
		var value int32
		if err := rows.Scan(&userID, &value); err != nil {
			return nil, err
		}
		res = append(res, model.Rating{
			UserId: model.UserId(userID),
			Value:  model.RatingValue(value),
		})
	}
	if len(res) == 0 {
		return nil, repository.ErrNotFound
	}
	return res, nil
}

func (r *Repository) Put(ctx context.Context, recordId model.RecordId, recordType model.RecordType, rating *model.Rating) error {
	query := `INSERT INTO ratings(record_id, record_type, user_id, value) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, recordId, recordType, rating.UserId, rating.Value)
	return err
}
