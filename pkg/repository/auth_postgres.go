package repository

import (
	"fmt"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user calendar.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING id;", usersTable)

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (calendar.User, error) {
	var user calendar.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2;", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetUserById(userId int) (calendar.User, error) {
	var user calendar.User
	query := fmt.Sprintf(`select us.username, us.email, us.is_active from %s us where us.id = $1;`, usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}
