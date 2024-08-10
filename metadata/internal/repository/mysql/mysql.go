package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-microservice/metadata/internal/repository"
	"go-microservice/metadata/pkg/model"
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

func (r *Repository) Get(ctx context.Context, id string) (*model.Metadata, error) {
	query := `SELECT title, description, director FROM movies WHERE id = ?`
	var title, description, director string
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&title, &description, &director); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &model.Metadata{Id: id, Title: title, Description: description, Director: director}, nil
}

func (r *Repository) Put(ctx context.Context, id string, metadata *model.Metadata) error {
	query := `INSERT INTO movies (id, title, description, director) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, id, metadata.Title, metadata.Description, metadata.Director)

	return err
}
