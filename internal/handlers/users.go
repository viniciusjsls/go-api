package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusjsls/go-api/internal/dto"
)

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @ID get-users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func UsersHandler(c *gin.Context) {
	users := []dto.UserResponse{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Create a new user
// @Description Create a new user with name
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User object"
// @Success 201 {object} User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var newUser dto.UserCreateRequest

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := dto.NewUserFromRequest(newUser)

	c.JSON(http.StatusCreated, dto.ToUserResponse(user))
}
