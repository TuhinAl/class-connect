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

	fmt.Println("==========PASS===========", student.Password)
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
		is_verified,
		class_id,
		class_name,
		password,
		father_phone,
		admission_fee,
		total_fee,
		remaining_fee,
		monthly_fee
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
	) RETURNING id, student_id, first_name, last_name, father_name, phone, gender,
	  course, email, is_active, is_verified, class_id, class_name, password, father_phone, admission_date, admission_fee, 
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
		student.IsVerfied,
		student.ClassID,
		student.ClassName,
		student.Password.Hash,
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
		&student.IsVerfied,
		&student.ClassID,
		&student.ClassName,
		&student.Password.Hash,
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

func (s *StudentStore) DeactivateStudentByID(ctx context.Context, studentReq *validation.StudentProxy) (*validation.StudentProxy, error) {

	query := `UPDATE student_information SET is_active = $2 WHERE id = $1 RETURNING id, first_name, last_name, student_id, email, is_active`

	var student validation.StudentProxy
	err := s.db.QueryRowContext(ctx, query, studentReq.Id, studentReq.IsActive).Scan(
		&student.Id,
		&student.FirstName,
		&student.LastName,
		&student.StudentId,
		&student.Email,
		&student.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that Id: %d", studentReq.Id)
		}
		return nil, fmt.Errorf("failed to update student status: %w", err)

	}
	return &student, err
}

func (s *StudentStore) GetAllStudents(ctx context.Context, limit int, offset int) ([]validation.StudentResponseProxy, int, error) {

	var response []validation.StudentResponseProxy
	totalResponse := 0
	query := `SELECT COUNT(*) OVER() AS total_rows, id, first_name, last_name, student_id,
	 email, phone from student_information WHERE is_active = true
	 LIMIT $1 OFFSET $2
	 `
	args := []interface{}{limit, offset}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var student validation.StudentResponseProxy
		err := rows.Scan(
			&totalResponse,
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.StudentId,
			&student.Email,
			&student.Phone,
		)
		if err != nil {
			return nil, 0, err
		}
		response = append(response, student)
	}
	return response, totalResponse, rows.Err()
}

func (s *StudentStore) GetStudentByEmail(ctx context.Context, email string) (*validation.StudentResponseProxy, error) {

	query := `SELECT id, first_name, last_name, student_id,
	 email, phone, password from student_information WHERE is_active = true and email = $1`

	var student validation.StudentResponseProxy
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&student.Id,
		&student.FirstName,
		&student.LastName,
		&student.StudentId,
		&student.Email,
		&student.Phone,
		&student.Password,
	)
	if err != nil {
		return nil, err
	}

	return &student, nil
}
