package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UsersHandler(c *gin.Context) {
	users := []user{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = rand.Intn(100)

	c.JSON(http.StatusCreated, newUser)
}
