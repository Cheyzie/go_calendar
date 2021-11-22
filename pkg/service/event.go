package service

import (
	calendar "github.com/cheyzie/go_calendar"
	"github.com/cheyzie/go_calendar/pkg/repository"
)

type EventService struct {
	repo         repository.Event
	calendarRepo repository.Calendar
}

func NewEventService(repo repository.Event, calendarRepo repository.Calendar) *EventService {
	return &EventService{repo: repo, calendarRepo: calendarRepo}
}

func (s *EventService) Create(calendarId, userId int, input calendar.Event) (int, error) {
	_, err := s.calendarRepo.GetById(calendarId, userId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(calendarId, input)
}

func (s *EventService) GetAll(calendarId, userId int) ([]calendar.Event, error) {
	return s.repo.GetAll(calendarId, userId)
}

func (s *EventService) GetById(eventId, userId int) (calendar.Event, error) {
	return s.repo.GetById(eventId, userId)
}

func (s *EventService) Delete(eventId, userId int) error {
	return s.repo.Delete(eventId, userId)
}

func (s *EventService) Update(eventId, userId int, input calendar.UpdateEventInput) error {
	return s.repo.Update(eventId, userId, input)
}
