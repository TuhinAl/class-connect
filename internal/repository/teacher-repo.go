package repository

import (
	"context"
	"database/sql"
	"golang-api/models"
)

type TeacherStore struct {
	db *sql.DB
}

type TeacherRepository interface {
	Create(ctx context.Context, teacher *models.Teacher) error
	// GetByID(ctx context.Context, id int64) (*models.Teacher, error)
	// GetAll(ctx context.Context) ([]models.Teacher, error)
	// Update(ctx context.Context, teacher *models.Teacher) error
	// Delete(ctx context.Context, id int64) error
	// GetByEmail(ctx context.Context, email string) (*models.Teacher, error)
	// GetByTeacherID(ctx context.Context, teacherID string) (*models.Teacher, error)
}

func (teacherStore *TeacherStore) Create(ctx context.Context, teacher *models.Teacher) error {
    return nil
}

