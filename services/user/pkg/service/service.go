package service

import (
	"user-service/pkg/models"
	"user-service/pkg/repository"

	"github.com/google/uuid"
)

type IUserService interface {
	CreateUser(models.User) error
	GetUserByE(email string) (models.User, error)
	GetUserByEP(email, password string) (models.User, error)
}

type IJwtManagerService interface {
	GeneratePairToken(userId uuid.UUID) (string, string, uuid.UUID, error)
	CompareTokens(refreshId uuid.UUID, token string) bool
	SaveRefreshToken(hashedToken string, userId, tokenId uuid.UUID) error
	DeleteRefreshTokenById(tokenId uuid.UUID) error
	GetRefreshTokenById(tokenId uuid.UUID) (string, error)
	ParseToken(accessToken string) (*models.AccessTokenClaims, error)
	CheckRefreshTokenExp(tokenId uuid.UUID) bool
}

type Service struct {
	IUserService
	IJwtManagerService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		IUserService:       NewUserService(repo.IUserRepo),
		IJwtManagerService: NewJwtManagerService(repo.IJwtManagerRepo),
	}
}
