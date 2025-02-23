package services

import (
	"gateway/pkg/models"

	"github.com/nats-io/nats.go"
)

type IUserService interface {
	CreateUser(user models.User) error
	GetUserByEP(email, passowrd string) (models.User, error)
	SignInByEP(email, passowrd string) (models.RefreshInput, error)
	SignInTokens(models.RefreshInput) (models.RefreshInput, error)
}

type Services struct {
	UserService IUserService
}

func NewService(ns *nats.Conn) *Services {
	return &Services{
		UserService: NewUserService(ns),
	}
}
