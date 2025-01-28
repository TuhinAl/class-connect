package validation

type StudentProxy struct {
	Id           int     `json:"id"`         // will Improve to uuid
	StudentId    int64   `json:"student_id"` // will Improve to uuid
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	FatherName   string  `json:"father_name"`
	Phone        string  `json:"phone"`
	Gender       string  `json:"gender"` // todo: enum
	Course       string  `json:"course"` // todo: enum
	Email        string  `json:"email"`
	IsActive     bool    `json:"is_active"`
	ClassID      int     `json:"class_id"`
	ClassName    string  `json:"class_name"` // todo: enum
	Password     string  `json:"password,omitempty"`
	FatherPhone  string  `json:"father_phone"`
	AdmissionFee float64 `json:"admission_fee"`
	TotalFee     float64 `json:"total_fee"`
	RemainingFee float64 `json:"remaining_fee"`
	MonthlyFee   float64 `json:"monthly_fee"`
}

type StudentResponseProxy struct {
	Id        int    `json:"id"`         // will Improve to uuid
	StudentId int64  `json:"student_id"` // will Improve to uuid
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
}

type StudentRequestProxy struct {
	Id            int    `json:"id"` // will Improve to uuid
	IsActive      bool   `json:"is_active"`
	AdmissionDate string `json:"admission_date"`
	Page          Page   `json:"page"`
}

type Page struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
