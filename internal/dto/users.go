package dto

import (
	"math/rand"

	"github.com/viniciusjsls/go-api/internal/models"
)

type UserCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewUserFromRequest(req UserCreateRequest) models.User {
	return models.User{
		ID:   rand.Intn(100),
		Name: req.Name,
	}
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}
