package repository

import (
	calendar "github.com/cheyzie/go_calendar"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(calendar.User) (int, error)
	GetUser(string, string) (calendar.User, error)
	GetUserById(int) (calendar.User, error)
}

type Calendar interface {
	Create(int, calendar.Calendar) (int, error)
	GetAll(int) ([]calendar.Calendar, error)
	GetById(int, int) (calendar.Calendar, error)
	Delete(int, int) error
	Update(int, int, calendar.UpdateCalendarInput) error
}

type Event interface {
	Create(int, calendar.Event) (int, error)
	GetAll(int, int) ([]calendar.Event, error)
	GetById(int, int) (calendar.Event, error)
	Delete(int, int) error
	Update(int, int, calendar.UpdateEventInput) error
}

type Repository struct {
	Authorization
	Calendar
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Calendar:      NewCalendarPostgres(db),
		Event:         NewEventPostgres(db),
	}
}
