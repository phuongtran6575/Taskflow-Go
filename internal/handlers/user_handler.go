package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/dto"
	"github.com/phuongtran6575/Taskflow-Go.git/internal/services/interfaces"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.Fail(c, http.StatusBadRequest, "MISSING_ID", "Missing user ID")
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		dto.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid user ID")
		return
	}

	user, err := h.userService.FindByID(uuidID)
	if err != nil {
		dto.Fail(c, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
		return
	}

	dto.OK(c, user)
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		dto.Fail(c, http.StatusInternalServerError, "FAILED_TO_FETCH_USERS", err.Error())
		return
	}

	dto.OK(c, users)
}
