package repository

import (
	"fmt"
	"user-service/pkg/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user models.User) error {
	query := fmt.Sprint(`
		INSERT INTO users (id, email, username, password) VALUES ($1, $2, $3, $4)
	`)
	_, err := r.db.Exec(query, user.Id, user.Email, user.Username, user.Password)
	return err
}

func (r *UserRepo) GetUserByE(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprint(`
		SELECT * FROM users WHERE email = $1
	`)
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *UserRepo) GetUserByEP(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprint(`
		SELECT * FROM users WHERE email = $1 and password = $2
	`)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
