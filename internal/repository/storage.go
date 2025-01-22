package repository

import (
	"database/sql"
)

type Storage struct {
	Student StudentRepository
	Teacher TeacherRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Student: &StudentStore{db: db},
		Teacher: &TeacherStore{db: db},
	}
}