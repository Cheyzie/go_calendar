package calendar

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active"`
}
