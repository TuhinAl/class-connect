package repository

import (
	"context"
	"database/sql"
	"golang-api/models"
)

type Storage struct {
	Student interface{
		CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	}
	//Teacher TeacherRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Student: &StudentStore{db: db},
		//Teacher: &TeacherStore{db: db},
	}
}