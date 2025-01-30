package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"golang-api/internal/validation"
	"golang-api/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
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
		// Password:      "hashed_password_here",
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
		// WriteJSONError(w, http.StatusBadRequest, err)
		app.BadRequestError(w, r, err)
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
		// Password:     payload.Password,
		FatherPhone:  payload.FatherPhone,
		AdmissionFee: payload.AdmissionFee,
		TotalFee:     payload.TotalFee,
		RemainingFee: payload.RemainingFee,
		MonthlyFee:   payload.MonthlyFee,
	}

	// createdStudent, err := repository.StudentRepository.CreateStudent(ctx, student)
	createdStudent, err := app.Store.Student.CreateStudent(ctx, student)

	if err != nil {
		// WriteJSONError(w, http.StatusInternalServerError, err)
		app.BadRequestError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdStudent)
}

func (app *ApplicationConfig) GetStudentByIdHandler(w http.ResponseWriter, r *http.Request) {

	sudentId := chi.URLParam(r, "studentId")
	id, err := strconv.ParseInt(sudentId, 10, 64)

	if err != nil {
		// WriteJSONError(w, http.StatusInternalServerError, err)
		app.InternalServerError(w, r, err)
		return
	}

	log.Println(id)
	fmt.Println("==========Student Id===========", id)

	ctx := r.Context()
	student, err := app.Store.Student.GetStudentByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.NotfoundError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)

}

func (app *ApplicationConfig) DeleteStudentByIdHandler(w http.ResponseWriter, r *http.Request) {

	sudentId := chi.URLParam(r, "studentId")
	id, err := strconv.ParseInt(sudentId, 10, 64)

	if err != nil {
		// WriteJSONError(w, http.StatusInternalServerError, err)
		app.InternalServerError(w, r, err)
		return
	}

	log.Println(id)
	fmt.Println("==========Student Id===========", id)

	ctx := r.Context()
	err = app.Store.Student.DeleteStudentByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.NotfoundError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(student)

}

func (app *ApplicationConfig) DeactivateStudentByIdHandler(w http.ResponseWriter, r *http.Request) {

	var req validation.StudentProxy

	if err := ReadJSONRequest(w, r, &req); err != nil {

		app.BadRequestError(w, r, err)
		return
	}

	log.Println(req.Id)
	fmt.Println("==========Student Id===========", req.Id)

	ctx := r.Context()
	response, err := app.Store.Student.DeactivateStudentByID(ctx, &req)

	// var proxyRespose *validation.StudentProxy

	proxyRespose := &validation.StudentResponseProxy{
		Id:        response.Id,
		StudentId: response.StudentId,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Phone:     response.Phone,
		Email:     response.Email,
		IsActive:  response.IsActive,
	}
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.NotfoundError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	WriteJSONResponse(w, http.StatusOK, &proxyRespose)

}

func (app *ApplicationConfig) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var req validation.StudentRequestProxy
	var custopResponse PageResponse
	var limit, offset int

	if err := ReadJSONRequest(w, r, &req); err != nil {

		app.BadRequestError(w, r, err)
		return
	}

	limit = req.Pageable.Size
	offset = req.Pageable.Page * limit

	ctx := r.Context()
	responses, totalRows, err := app.Store.Student.GetAllStudents(ctx, limit, offset)

	// var proxyRespose *validation.StudentProxy
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.NotfoundError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}
	custopResponse = PageResponse{
		TotalRows: totalRows,
		CurrPage:  req.Pageable.Page,
		TotalPage: (totalRows / limit),
		Data:      responses,
	}

	WriteJSONResponse(w, http.StatusOK, custopResponse)

}


func (app *ApplicationConfig) GetStudentByEmailHandler(w http.ResponseWriter, r *http.Request) {

	var req validation.StudentRequestProxy

	if err := ReadJSONRequest(w, r, &req); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	responses, err := app.Store.Student.GetStudentByEmail(ctx, req.Email)

	// var proxyRespose *validation.StudentProxy
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.NotfoundError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}
	WriteJSONResponse(w, http.StatusOK, &responses)

}
