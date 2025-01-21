package models

type Student struct {
	ID   int
	StudentId   int64 // will Improve to uuid
	FirstName string
	LastName string
	FatherName string
	Phone string
	Gender string // todo: enum
	Course string // todo: enum
	Email string
	IsActive bool
	ClassID int
	Class string // todo: enum
	Password string
	FatherPhone string
	AdmissionDate string // todo: date
	AdmissionFee float64
	TotalFee float64
	RemainingFee float64
	MonthlyFee float64
}