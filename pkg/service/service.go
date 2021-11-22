package service

import (
	calendar "github.com/cheyzie/go_calendar"
	"github.com/cheyzie/go_calendar/pkg/repository"
)

type Authorization interface {
	CreateUser(calendar.User) (int, error)
	GenerateToken(string, string) (calendar.Credentionals, error)
	ParseToken(string) (int, error)
	GetUserData(int) (calendar.User, error)
}

type Calendar interface {
	Create(int, calendar.Calendar) (int, error)
	GetAll(int) ([]calendar.Calendar, error)
	GetById(int, int) (calendar.Calendar, error)
	Delete(int, int) error
	Update(int, int, calendar.UpdateCalendarInput) error
}

type Event interface {
	Create(int, int, calendar.Event) (int, error)
	GetAll(int, int) ([]calendar.Event, error)
	GetById(int, int) (calendar.Event, error)
	Delete(int, int) error
	Update(int, int, calendar.UpdateEventInput) error
}

type Service struct {
	Authorization
	Calendar
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Calendar:      NewCalendarService(repos.Calendar),
		Event:         NewEventService(repos.Event, repos.Calendar),
	}
}
