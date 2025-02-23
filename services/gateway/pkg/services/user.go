package services

import (
	"encoding/json"
	"errors"
	"gateway/pkg/models"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type UserService struct {
	ns *nats.Conn
}

func NewUserService(ns *nats.Conn) *UserService {
	return &UserService{ns: ns}
}

func (s *UserService) CreateUser(user models.User) error {
	user.Id = uuid.New()
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	msg, err := s.ns.Request("user_service.create_user", data, time.Second*10)
	if err != nil {
		return err
	}
	if string(msg.Data) == "success" {
		return nil
	} else {
		return errors.New(string(msg.Data))
	}

}
func (s *UserService) GetUserByEP(email, passowrd string) (models.User, error) {
	return models.User{}, nil
}

func (s *UserService) SignInByEP(email, passowrd string) (models.RefreshInput, error) {
	data, err := json.Marshal(models.User{Email: email, Password: passowrd})
	if err != nil {
		return models.RefreshInput{}, err
	}
	msg, err := s.ns.Request("user_service.get_tokens", data, time.Second*10)
	if err != nil {
		return models.RefreshInput{}, err
	}
	var tokens models.RefreshInput
	err = json.Unmarshal(msg.Data, &tokens)
	return tokens, err
}
func (s *UserService) SignInTokens(models.RefreshInput) (models.RefreshInput, error) {
	return models.RefreshInput{}, nil
}
