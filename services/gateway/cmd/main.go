package main

import (
	"gateway/pkg/server"
	"gateway/pkg/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

/*
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
*/

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while loading .env file: %s", err.Error())
	}

	ns, err := nats.Connect(os.Getenv("NATSURL"))
	if err != nil {
		log.Fatalf("Error while connect to NATS: %s", err.Error())

	}
	services := services.NewService(ns)

	srv := server.NewServer(*services)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	srv.Run(port)

}
