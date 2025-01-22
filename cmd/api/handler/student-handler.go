package handler

import (
	"encoding/json"
	"net/http"
	"golang-api/models"
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
        Class:         "Grade 12",
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
