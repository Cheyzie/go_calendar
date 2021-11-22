package service

import (
	calendar "github.com/cheyzie/go_calendar"
	"github.com/cheyzie/go_calendar/pkg/repository"
)

type CalendarService struct {
	repo repository.Calendar
}

func NewCalendarService(repo repository.Calendar) *CalendarService {
	return &CalendarService{repo: repo}
}

func (s *CalendarService) Create(userId int, input calendar.Calendar) (int, error) {
	return s.repo.Create(userId, input)
}

func (s *CalendarService) GetAll(userId int) ([]calendar.Calendar, error) {
	return s.repo.GetAll(userId)
}

func (s *CalendarService) GetById(calendarId, userId int) (calendar.Calendar, error) {
	return s.repo.GetById(calendarId, userId)
}

func (s *CalendarService) Delete(calendaId, userId int) error {
	return s.repo.Delete(calendaId, userId)
}

func (s *CalendarService) Update(calendarId, userId int, input calendar.UpdateCalendarInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(calendarId, userId, input)
}
