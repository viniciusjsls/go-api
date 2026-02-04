package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @ID get-users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func UsersHandler(c *gin.Context) {
	users := []User{
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
	var newUser User
	c.BindJSON(&newUser)

	newUser.ID = rand.Intn(100)

	c.JSON(http.StatusCreated, newUser)
}
