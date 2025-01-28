package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-api/internal/validation"
	"golang-api/models"
)

var (
	ErrorStudentNotFound error = errors.New("student not found")
)

type Storage struct {
	Student interface {
		CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
		GetStudentByID(ctx context.Context, id int64) (*models.Student, error)
		DeleteStudentByID(ctx context.Context, id int64) error
		DeactivateStudentByID(ctx context.Context, student *validation.StudentProxy) (*validation.StudentProxy, error)
		GetAllStudents(ctx context.Context) ([]validation.StudentResponseProxy, error)
	}
	//Teacher TeacherRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Student: &StudentStore{db: db},
		//Teacher: &TeacherStore{db: db},
	}
}
