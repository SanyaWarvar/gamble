package worker

import (
	"encoding/json"
	"log"
	"user-service/pkg/models"
	"user-service/pkg/service"

	"github.com/nats-io/nats.go"
)

type Worker struct {
	ns *nats.Conn
	s  *service.Service
}

func NewWorker(ns *nats.Conn, s *service.Service) *Worker {
	return &Worker{
		ns: ns,
		s:  s,
	}
}

func (w *Worker) Run() {
	subjectPrefix := "user_service."

	subject := subjectPrefix + "create_user"
	_, err := w.ns.Subscribe(subject, func(m *nats.Msg) {
		var user models.User
		data := m.Data
		err := json.Unmarshal(data, &user)
		if err != nil {
			log.Printf("Error while unmarshal json: %s", err.Error())
			m.Respond([]byte(err.Error()))
			return
		}
		err = w.s.IUserService.CreateUser(user)
		if err != nil {
			log.Printf("Error while insert into db: %s", err.Error())
			m.Respond([]byte(err.Error()))
			return
		}
		m.Respond([]byte("success"))
	})
	if err != nil {
		log.Print(err.Error())
	}

	subject = subjectPrefix + "sign_in_ep"
	_, err = w.ns.Subscribe(subject, func(m *nats.Msg) {
		var user models.User
		data := m.Data
		err := json.Unmarshal(data, &user)
		if err != nil {
			log.Printf("Error while unmarshal json: %s", err.Error())
			m.Respond([]byte(err.Error()))
			return
		}
		err = w.s.IUserService.CreateUser(user)
		if err != nil {
			log.Printf("Error while insert into db: %s", err.Error())
			m.Respond([]byte(err.Error()))
			return
		}

	})
	if err != nil {
		log.Print(err.Error())
	}

	subject = subjectPrefix + "get_tokens"
	_, err = w.ns.Subscribe(subject, func(m *nats.Msg) {
		var user models.User
		data := m.Data
		err := json.Unmarshal(data, &user)
		if err != nil {
			log.Printf("Error while unmarshal json: %s", err.Error())
			return
		}

		target, err := w.s.IUserService.GetUserByEP(user.Email, user.Password)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		token, refresh, _, err := w.s.IJwtManagerService.GeneratePairToken(target.Id)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}
		dataToSend, err := json.Marshal(map[string]string{"access_token": token, "refresh_token": refresh})
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}
		err = m.Respond(dataToSend)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}
	})
	if err != nil {
		log.Print(err.Error())
	}

	select {}
}
