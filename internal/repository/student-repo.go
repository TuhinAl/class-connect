package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-api/internal/validation"
	"golang-api/models"
)

type StudentStore struct {
	db *sql.DB
}

/* type StudentRepository interface {

	// GetByID(ctx context.Context, id int64) (*models.Student, error)
	// GetAll(ctx context.Context) ([]models.Student, error)
	// Update(ctx context.Context, student *models.Student) error
	// Delete(ctx context.Context, id int64) error
	// GetByEmail(ctx context.Context, email string) (*models.Student, error)
	// GetByStudentID(ctx context.Context, studentID int64) (*models.Student, error)
} */

func (studentStore *StudentStore) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {

	query := `INSERT INTO student_information (
		student_id,
		first_name,
		last_name,
		father_name,
		phone,
		gender,
		course,
		email,
		is_active,
		class_id,
		class_name,
		password,
		father_phone,
		admission_fee,
		total_fee,
		remaining_fee,
		monthly_fee
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
	) RETURNING id, student_id, first_name, last_name, father_name, phone, gender,
	  course, email, is_active, class_id, class_name, password, father_phone, admission_date, admission_fee, 
	  total_fee, remaining_fee, monthly_fee
	`
	err := studentStore.db.QueryRowContext(
		ctx,
		query,
		student.StudentId,
		student.FirstName,
		student.LastName,
		student.FatherName,
		student.Phone,
		student.Gender,
		student.Course,
		student.Email,
		student.IsActive,
		student.ClassID,
		student.ClassName,
		student.Password,
		student.FatherPhone,
		student.AdmissionFee,
		student.TotalFee,
		student.RemainingFee,
		student.MonthlyFee,
	).Scan(
		&student.ID,
		&student.StudentId,
		&student.FirstName,
		&student.LastName,
		&student.FatherName,
		&student.Phone,
		&student.Gender,
		&student.Course,
		&student.Email,
		&student.IsActive,
		&student.ClassID,
		&student.ClassName,
		&student.Password,
		&student.FatherPhone,
		&student.AdmissionDate,
		&student.AdmissionFee,
		&student.TotalFee,
		&student.RemainingFee,
		&student.MonthlyFee,
	)

	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentStore) GetStudentByID(ctx context.Context, id int64) (*models.Student, error) {

	query := `SELECT id, student_id, first_name, last_name, father_name, phone, gender,
	  course, email, is_active, class_id, class_name, password, father_phone, admission_date, admission_fee, 
	  total_fee, remaining_fee, monthly_fee
	FROM student_information WHERE id = $1`

	var student models.Student
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&student.ID,
		&student.StudentId,
		&student.FirstName,
		&student.LastName,
		&student.FatherName,
		&student.Phone,
		&student.Gender,
		&student.Course,
		&student.Email,
		&student.IsActive,
		&student.ClassID,
		&student.ClassName,
		&student.Password,
		&student.FatherPhone,
		&student.AdmissionDate,
		&student.AdmissionFee,
		&student.TotalFee,
		&student.RemainingFee,
		&student.MonthlyFee,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorStudentNotFound
		default:
			return nil, err
		}
	}
	return &student, nil
}

func (s *StudentStore) DeleteStudentByID(ctx context.Context, id int64) error {

	query := `DELETE FROM student_information WHERE id = $1`

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err

	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err

	}
	if rows == 0 {
		return ErrorStudentNotFound
	}
	return nil
}

func (s *StudentStore) DeactivateStudentByID(ctx context.Context, id int64, activeStatus bool) (*validation.StudentProxy, error) {

	query := `UPDATE student_information SET is_active = $2 WHERE id = $1 RETURNING id, first_name, last_name, student_id, email, is_active`

	var student validation.StudentProxy
	err := s.db.QueryRowContext(ctx, query, id, activeStatus).Scan(
		&student.Id,
		&student.FirstName,
		&student.LastName,
		&student.StudentId,
		&student.Email,
		&student.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that Id: %d", id)
		}
		return nil, fmt.Errorf("failed to update student status: %w", err)

	}
	return &student, err
}
