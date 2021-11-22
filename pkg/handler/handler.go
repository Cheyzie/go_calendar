package handler

import (
	"time"

	"github.com/cheyzie/go_calendar/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		cors.New(cors.Config{
			AllowAllOrigins:        true,
			AllowMethods:           []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:           []string{"ORIGIN", "Authorization", "Content-Type"},
			AllowCredentials:       true,
			AllowBrowserExtensions: true,
			MaxAge:                 300 * time.Second,
		}),
	)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/me", h.getUserData)
		calendars := api.Group("/calendars")
		{
			calendars.POST("/", h.createCalendar)
			calendars.GET("/", h.getAllCalendars)
			calendars.GET("/:id", h.getCalendarById)
			calendars.PUT("/:id", h.updateCalendar)
			calendars.DELETE("/:id", h.deleteCalendar)

			// subscribes := calendars.Group(":id/subscribes")
			// {
			// 	subscribes.POST("/", h.AddSubscribe)
			// 	subscribes.GET("/", h.GetAllSubscribes)
			// 	subscribes.PUT("/:subscribe_id", h.UpdateSubscribe)
			// 	subscribes.DELETE("/:subscribe_id", h.DeleteSubscribe)
			// }

			events := calendars.Group(":id/events")
			{
				events.POST("/", h.createEvent)
				events.GET("/", h.getAllEvents)
			}
		}
		events := api.Group("events")
		{
			events.GET("/:event_id", h.getEventById)
			events.PUT("/:event_id", h.updateEvent)
			events.DELETE("/:event_id", h.deleteEvent)
		}
	}
	return router
}
