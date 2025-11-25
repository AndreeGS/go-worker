package main

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Erro ao abrir canal:", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"tasks",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Erro ao consumir fila:", err)
	}

	fmt.Println("Worker iniciado. Aguardando mensagens...")

	for msg := range msgs {
		//Processamento
		fmt.Println("Processando tarefa:", string(msg.Body))
	}
}
