package rabbitmq

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Flikest/myMicroservices/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Send(message string) {
	connectingString := fmt.Sprintf("amqp://%s:%s@localhost:5673/", os.Getenv("MQ_USER"), os.Getenv("MQ_PASS"))

	conn, err := amqp.Dial(connectingString)
	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	chanel, err := conn.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	defer chanel.Close()

	err = chanel.ExchangeDeclare(
		"users",  // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	errors.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := []byte(message)
	err = chanel.PublishWithContext(ctx,
		"users", // exchange
		"",      // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	errors.FailOnError(err, "Failed to publish a message")

	slog.Info("Send: %s", body)

}
