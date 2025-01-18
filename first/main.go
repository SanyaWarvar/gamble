package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	subject := "updates"
	fmt.Println("Запуск \"отправителя\"")
	for {
		msg := faker.Word()
		data, _ := json.Marshal(map[string]string{"message": msg})
		fmt.Printf("Отправлено: %v\n", data)
		nc.Publish(subject, data)
		time.Sleep(5 * time.Second)
	}
}
