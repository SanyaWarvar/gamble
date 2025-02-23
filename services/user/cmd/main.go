package main

import (
	"log"
	"os"
	"time"
	"user-service/pkg/repository"
	"user-service/pkg/service"
	"user-service/pkg/worker"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

func main() {
	time.Sleep(time.Second * 5)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while loading .env file: %s", err.Error())
	}

	ns, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer ns.Close()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Error while create connection to db: %s", err.Error())
	}

	accessTokenTTL, err := time.ParseDuration(os.Getenv("ACCESSTOKENTTL"))
	if err != nil {
		log.Fatalf("Errof while parse accessTokenTTL: %s", err.Error())
	}
	refreshTokenTTL, err := time.ParseDuration(os.Getenv("REFRESHTOKENTTL"))
	if err != nil {
		log.Fatalf("Errof while parse refreshTokenTTL: %s", err.Error())
	}
	jwtCfg := repository.NewJwtManagerCfg(accessTokenTTL, refreshTokenTTL, os.Getenv("SIGNINGKEY"), jwt.SigningMethodHS256)

	r := repository.NewRepository(db, jwtCfg)
	s := service.NewService(*r)
	w := worker.NewWorker(ns, s)

	w.Run()
}
