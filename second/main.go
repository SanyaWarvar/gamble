package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	fmt.Println("Запуск \"получателя\"")

	subject := "updates"
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Printf("Получено сообщение: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
