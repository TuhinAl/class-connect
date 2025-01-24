package handler

import (
	"encoding/json"
	"golang-api/internal/validation"
	"golang-api/models"
	"net/http"
	"time"
)

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	student := models.Student{
		ID:            1,
		StudentId:     1001,
		FirstName:     "John",
		LastName:      "Doe",
		FatherName:    "Michael Doe",
		Phone:         "+880-1234567890",
		Gender:        "Male",
		Course:        "Computer Science",
		Email:         "john.doe@example.com",
		IsActive:      true,
		ClassID:       101,
		ClassName:     "Grade 12",
		Password:      "hashed_password_here",
		FatherPhone:   "+880-9876543210",
		AdmissionDate: time.Now().Format("2006-01-02"),
		AdmissionFee:  5000.00,
		TotalFee:      50000.00,
		RemainingFee:  45000.00,
		MonthlyFee:    5000.00,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (app *ApplicationConfig) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {

	var payload validation.StudentProxy
	if err := ReadJSONRequest(w, r, &payload); err != nil {
		WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	// Create the student
	student := &models.Student{
		StudentId:    payload.StudentId,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		FatherName:   payload.FatherName,
		Phone:        payload.Phone,
		Gender:       payload.Gender,
		Course:       payload.Course,
		Email:        payload.Email,
		IsActive:     payload.IsActive,
		ClassID:      payload.ClassID,
		ClassName:    payload.ClassName,
		Password:     payload.Password,
		FatherPhone:  payload.FatherPhone,
		AdmissionFee: payload.AdmissionFee,
		TotalFee:     payload.TotalFee,
		RemainingFee: payload.RemainingFee,
		MonthlyFee:   payload.MonthlyFee,
	}

	// createdStudent, err := repository.StudentRepository.CreateStudent(ctx, student)
	createdStudent, err := app.Store.Student.CreateStudent(ctx, student)

	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdStudent)
}
