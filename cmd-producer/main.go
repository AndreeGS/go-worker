package main

import (
	"log"

	"github.com/gin-gonic/gin"
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

	queue, err := ch.QueueDeclare(
		"tasks",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Erro ao declarar fila:", err)
	}

	r := gin.Default()

	r.POST("/task", func(c *gin.Context) {
		type Input struct {
			Message string `json:"message"`
		}

		var input Input
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "JSON inválido"})
			return
		}

		err = ch.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(input.Message),
			},
		)

		if err != nil {
			c.JSON(500, gin.H{"error": "Erro ao enviar mensagem"})
			return
		}

		c.JSON(200, gin.H{"status": "Mensagem enviada!"})
	})

	r.Run(":8080")
}
