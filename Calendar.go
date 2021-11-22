package calendar

import (
	"errors"
	"time"
)

type Calendar struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Type        string `json:"type" db:"type"`
}

type Subscribe struct {
	Id         int
	UserId     int
	CalendarId int
	Type       string
}

type Event struct {
	Id          int       `json:"id" db:"id"`
	CalendarId  int       `json:"-"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	Time        time.Time `json:"time" db:"time" binding:"required"`
}

type UpdateCalendarInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i *UpdateCalendarInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("input is empty")
	}
	return nil
}

type UpdateEventInput struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Time        *time.Time `json:"time"`
}

func (i *UpdateEventInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Time != nil {
		return errors.New("input is empty")
	}
	return nil
}
