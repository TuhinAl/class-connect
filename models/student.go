package models

import (
	"errors"
	"golang-api/internal/utility/token"

	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	ID            int        `json:"id"`
	StudentId     int64      `json:"student_id"` // will Improve to uuid
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	FatherName    string     `json:"father_name"`
	Phone         string     `json:"phone"`
	Gender        string     `json:"gender"` // todo: enum
	Course        string     `json:"course"` // todo: enum
	Email         string     `json:"email"`
	IsActive      bool       `json:"is_active"`
	IsVerfied     bool       `json:"is_verified"`
	ClassID       int        `json:"class_id"`
	ClassName     string     `json:"class_name"` // todo: enum
	Password      password   `json:"-"`
	FatherPhone   string     `json:"father_phone"`
	AdmissionDate string     `json:"admission_date"` // todo: date
	AdmissionFee  float64    `json:"admission_fee"`
	TotalFee      float64    `json:"total_fee"`
	RemainingFee  float64    `json:"remaining_fee"`
	MonthlyFee    float64    `json:"monthly_fee"`
	Role          token.Role `json:"role"`
}

// plaintext and hashed versions of the password for a user
type password struct {
	Plaintext *string
	Hash      []byte
}

// calculates the bcrypt hash of a plaintext password, and stores both the hash and the plaintext versions in the struct
func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

/* checks whether the provided plaintext password matches the
hashed password stored in the struct, returning true if it matches and false otherwise.
*/

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
