package handler

import (
	"net/http"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/gin-gonic/gin"
)

// @Sumary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body calendar.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [POST]
func (h *Handler) signUp(c *gin.Context) {
	var input calendar.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Sumary SignIn
// @Tags auth
// @Description Log in
// @ID logIn
// @Accept json
// @Produce json
// @Param input body signInInput true "Data for logIn"
// @Success 200 {object} calendar.Credentionals
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [POST]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	credentionals, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, credentionals)
}

// @Sumary GetUserData
// @Tags user
// @Description get youself info
// @ID get-user-data
// @Accept json
// @Produce json
// @Param input body
// @Success 200 {object} calendar.User
// @Failure 400, 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/me [POST]
func (h *Handler) getUserData(c *gin.Context) {
	ctxId, _ := c.Get(userCtx)
	id, _ := ctxId.(int)

	user, err := h.services.Authorization.GetUserData(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
