package service

import (
	"errors"
	"user-service/pkg/models"
	"user-service/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByE(email string) (models.User, error) {
	return s.repo.GetUserByE(email)
}

func (s *UserService) GetUserByEP(email, password string) (models.User, error) {
	var user models.User
	targetUser, err := s.repo.GetUserByE(email)
	if err != nil {
		return user, err
	}

	if bcrypt.CompareHashAndPassword([]byte(targetUser.Password), []byte(password)) == nil {
		return targetUser, err
	}
	return user, errors.New("incorrect password")
}
