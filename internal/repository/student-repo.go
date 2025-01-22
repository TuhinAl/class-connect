package repository

import (
	"context"
	"database/sql"
	"golang-api/models"
)

type StudentStore struct {
	db *sql.DB
}

type StudentRepository interface {
	Create(ctx context.Context, student *models.Student) error
	// GetByID(ctx context.Context, id int64) (*models.Student, error)
	// GetAll(ctx context.Context) ([]models.Student, error)
	// Update(ctx context.Context, student *models.Student) error
	// Delete(ctx context.Context, id int64) error
	// GetByEmail(ctx context.Context, email string) (*models.Student, error)
	// GetByStudentID(ctx context.Context, studentID int64) (*models.Student, error)
}

func (studentStore *StudentStore) Create(ctx context.Context, student *models.Student) error {
	return nil
}