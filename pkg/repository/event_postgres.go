package repository

import (
	"fmt"
	"strings"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/jmoiron/sqlx"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) Create(calendarId int, event calendar.Event) (int, error) {

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (calendar_id, title, description, time) values ($1, $2, $3, $4) RETURNING id", eventsTable)

	row := r.db.QueryRow(createItemQuery, calendarId, event.Title, event.Description, event.Time)
	err := row.Scan(&itemId)

	return itemId, err
}

func (r *EventPostgres) GetAll(calendarId, userId int) ([]calendar.Event, error) {
	var events []calendar.Event
	query := fmt.Sprintf(
		`select ev.id, ev.title, ev.description, ev.time from %s ev left join %s cl
		on ev.calendar_id = cl.id left join %s sb on cl.id = sb.calendar_id where cl.id = $1 and sb.user_id = $2;`,
		eventsTable, calendarsTable, subscribesTable)
	if err := r.db.Select(&events, query, calendarId, userId); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventPostgres) GetById(eventId, userId int) (calendar.Event, error) {
	var item calendar.Event
	query := fmt.Sprintf(
		`select ev.id, ev.title, ev.description, ev.time from %s ev left join %s cl
		on ev.calendar_id = cl.id left join %s sb on cl.id = sb.calendar_id where ev.id = $1 and sb.user_id = $2;`,
		eventsTable, calendarsTable, subscribesTable)
	if err := r.db.Get(&item, query, eventId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *EventPostgres) Delete(eventId, userId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s ev USING %s cl, %s sb 
		WHERE ev.calendar_id = cl.id AND cl.id = sb.calendar_id AND sb.user_id = $1 AND ev.id = $2;`,
		eventsTable, calendarsTable, subscribesTable)
	_, err := r.db.Exec(query, userId, eventId)
	return err
}

func (r *EventPostgres) Update(eventId, userId int, input calendar.UpdateEventInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Time != nil {
		setValues = append(setValues, fmt.Sprintf("time=$%d", argId))
		args = append(args, *input.Time)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s ev SET %s FROM %s cl, %s sb
		WHERE ev.calendar_id = cl.id AND sb.calendar_id = cl.id AND sb.user_id = $%d AND ev.id = $%d`,
		eventsTable, setQuery, calendarsTable, subscribesTable, argId, argId+1)
	args = append(args, userId, eventId)

	_, err := r.db.Exec(query, args...)
	return err
}
