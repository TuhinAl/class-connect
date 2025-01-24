package validation

type StudentProxy struct {
	StudentId     int64   `json:"student_id"` // will Improve to uuid
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	FatherName    string  `json:"father_name"`
	Phone         string  `json:"phone"`
	Gender        string  `json:"gender"` // todo: enum
	Course        string  `json:"course"` // todo: enum
	Email         string  `json:"email"`
	IsActive      bool    `json:"is_active"`
	ClassID       int     `json:"class_id"`
	ClassName     string  `json:"class_name"` // todo: enum
	Password      string  `json:"password,omitempty"`
	FatherPhone   string  `json:"father_phone"`
	AdmissionFee  float64 `json:"admission_fee"`
	TotalFee      float64 `json:"total_fee"`
	RemainingFee  float64 `json:"remaining_fee"`
	MonthlyFee    float64 `json:"monthly_fee"`
}