package handler

import (
	"net/http"
	"strconv"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/gin-gonic/gin"
)

// @Sumary CreateEvent
// @Tags events
// @Description Create event
// @ID create-event
// @Accept json
// @Produce json
// @Param input calendar_id path int true "Calendar ID"
// @Param body {object} calendar.Event
// @Success 200 {string} str 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendars/{calendar_id}/events [POST]
func (h *Handler) createEvent(c *gin.Context) {
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

	var input calendar.Event

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Event.Create(calendarId, userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type getAllEventsResponse struct {
	Events []calendar.Event
}

// @Sumary GetAllEvent
// @Tags events
// @Description Get all event
// @ID get-all-events
// @Produce json
// @Param input calendar_id path int true "Calendar ID"
// @Success 200 {object} getAllEventsResponse
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/calendar/{calendar_id}/events/ [GET]
func (h *Handler) getAllEvents(c *gin.Context) {
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
	events, err := h.services.Event.GetAll(calendarId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllEventsResponse{Events: events})
}

// @Sumary GetEventByID
// @Tags events
// @Description Get event by ID
// @ID get-event-by-id
// @Accept json
// @Produce json
// @Param input event_id path int true "Event ID"
// @Success 200 {object} calendar.Event
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/{event_id} [GET]
func (h *Handler) getEventById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	event, err := h.services.Event.GetById(eventId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, event)
}

// @Sumary UpdateEvent
// @Tags events
// @Description Update event
// @ID update-event
// @Accept json
// @Produce json
// @Param input event_id path int true "Event ID"
// @Param body {object} calendar.UpdateEventInput
// @Success 200 {string} str 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/{event_id} [PUT]
func (h *Handler) updateEvent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input calendar.UpdateEventInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Event.Update(eventId, userId, input); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Sumary DeleteEvent
// @Tags events
// @Description Delete event
// @ID delete-event
// @Produce json
// @Param input event_id path int true "Event ID"
// @Success 200 {string} str 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/events/{event_id} [DELETE]
func (h *Handler) deleteEvent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	eventId, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Event.Delete(eventId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
