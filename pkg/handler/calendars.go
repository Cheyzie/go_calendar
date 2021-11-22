package handler

import (
	"net/http"
	"strconv"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/gin-gonic/gin"
)

// @Sumary CreateCalendar
// @Tags calendars
// @Description Create new calendar
// @ID create-calendar
// @Accept json
// @Produce json
// @Param input body calendar.Calendar true "Data for logIn"
// @Success 200 {integer} int 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendars/ [POST]
func (h *Handler) createCalendar(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input calendar.Calendar

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Calendar.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type getAllCalendarsResponse struct {
	Calendars []calendar.Calendar
}

// @Sumary CreateCalendar
// @Tags calendars
// @Description Create new calendar
// @ID create-calendar
// @Produce json
// @Success 200 {object} getAllCalendarsResponse
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendars/ [POST]
func (h *Handler) getAllCalendars(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	calendars, err := h.services.Calendar.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllCalendarsResponse{
		Calendars: calendars,
	})
}

// @Sumary GetCalendarById
// @Tags calendars
// @Description Get calendar by ID
// @ID get-calendar-by-id
// @Produce json
// @Param input calendar_id path int true "Calendar ID"
// @Success 200 {integer} int 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendars/{calendar_id} [GET]
func (h *Handler) getCalendarById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	calendarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	calendar, err := h.services.Calendar.GetById(calendarId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, calendar)
}

// @Sumary UpdateCalendar
// @Tags calendars
// @Description Update calendar
// @ID update-calendar
// @Accept json
// @Produce json
// @Param input calendar_id path int true "Calendar ID"
// @Param body {object} calendar.UpdateCalendarInput
// @Success 200 {string} str 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendars/{calendar_id} [PUT]
func (h *Handler) updateCalendar(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	calendarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input calendar.UpdateCalendarInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Calendar.Update(calendarId, userId, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Sumary DeleteCalendar
// @Tags calendars
// @Description Delete calendar
// @ID delete-calendar
// @Produce json
// @Param input calendar_id path int true "Calendar ID"
// @Success 200 {string} str 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendars/{calendar_id} [DELETE]
func (h *Handler) deleteCalendar(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	calendarId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	err = h.services.Calendar.Delete(calendarId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
