package repository

import (
	"user-service/pkg/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	CreateUser(models.User) error
	GetUserByE(email string) (models.User, error)
	GetUserByEP(email, password string) (models.User, error)
}

type IJwtManagerRepo interface {
	GenerateAccessToken(userId, refreshId uuid.UUID) (string, error)
	GenerateRefreshToken(userId uuid.UUID) (string, error)
	GeneratePairToken(userId uuid.UUID) (string, string, uuid.UUID, error)
	CompareTokens(hashedToken, token string) bool
	HashToken(refreshToken string) (string, error)
	SaveRefreshToken(hashedToken string, tokenId, userId uuid.UUID) error
	DeleteRefreshTokenById(tokenId uuid.UUID) error
	GetRefreshTokenById(tokenId uuid.UUID) (string, error)
	ParseToken(accessToken string) (*models.AccessTokenClaims, error)
	CheckRefreshTokenExp(tokenId uuid.UUID) bool
}

type Repository struct {
	IUserRepo
	IJwtManagerRepo
}

func NewRepository(db *sqlx.DB, cfg *JwtManagerCfg) *Repository {
	return &Repository{
		IUserRepo:       NewUserRepo(db),
		IJwtManagerRepo: NewJwtManagerPostgres(db, cfg),
	}
}
