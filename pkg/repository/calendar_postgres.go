package repository

import (
	"fmt"
	"strings"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type calendarPostgres struct {
	db *sqlx.DB
}

func NewCalendarPostgres(db *sqlx.DB) *calendarPostgres {
	return &calendarPostgres{db: db}
}

func (r *calendarPostgres) Create(userId int, input calendar.Calendar) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", calendarsTable)
	row := tx.QueryRow(createListQuery, input.Title, input.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createSubscribeQuery := fmt.Sprintf("INSERT INTO %s (user_id, calendar_id, type) VALUES ($1, $2, 'creator')", subscribesTable)
	_, err = tx.Exec(createSubscribeQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *calendarPostgres) GetAll(userId int) ([]calendar.Calendar, error) {
	var lists []calendar.Calendar

	query := fmt.Sprintf("SELECT cl.id, cl.title, cl.description, sb.type FROM %s cl RIGHT JOIN %s sb on cl.id = sb.calendar_id WHERE sb.user_id = $1",
		calendarsTable, subscribesTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *calendarPostgres) GetById(calendarId, userId int) (calendar.Calendar, error) {
	var calendarInfo calendar.Calendar

	query := fmt.Sprintf(`SELECT cl.id, cl.title, cl.description, sb.type FROM %s cl
		INNER JOIN %s sb on cl.id = sb.calendar_id WHERE sb.user_id = $1 AND cl.id = $2;`,
		calendarsTable, subscribesTable)
	err := r.db.Get(&calendarInfo, query, userId, calendarId)

	return calendarInfo, err
}

func (r *calendarPostgres) Delete(calendarId, userId int) error {
	query := fmt.Sprintf(
		"delete from %s cl using %s sb where cl.id = sb.calendar_id AND sb.user_id=$1 AND cl.id=$2 and sb.type = 'creator';",
		calendarsTable, subscribesTable)
	_, err := r.db.Exec(query, userId, calendarId)

	return err
}

func (r *calendarPostgres) Update(calendarId, userId int, input calendar.UpdateCalendarInput) error {
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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"update %s cl set %s from %s sb where cl.id = sb.calendar_id and cl.id=$%d and sb.user_id=$%d and sb.type = 'creator';",
		calendarsTable, setQuery, subscribesTable, argId, argId+1)
	args = append(args, calendarId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
