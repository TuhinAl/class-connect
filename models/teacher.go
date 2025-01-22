package models

import "time"

type Teacher struct {
	Id           int64     `json:"id"`
	TeacherId    string    `json:"teacherId"`
	JoinDate     time.Time `json:"joinDate"`
	Name         string    `json:"name"`
	TakenSubject string    `json:"takenSubject"`
	IsActive     bool      `json:"isActive"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
	Password     string    `json:"password,omitempty"`
}