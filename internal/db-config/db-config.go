package dbconfig



type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// func NewDB(db *sql.DB) repository.Storage {
// 	return repository.Storage{
// 		Student: &repository.StudentStore{db},
// 		Teacher: &repository.TeacherStore{db},
// 	}
// }