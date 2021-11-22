package calendar

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty"`
	IsActive bool   `json:"is_active" db:"is_active"`
}
