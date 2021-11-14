package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Config ...
type Config struct {
	Conn          *amqp.Connection
	ExchangeName  string
	QueueName     string
	PrefetchCount int
	VHost         string
}
/*

// Send ...
func (rs *Config) Send(customer *models.Customer, uuid uuid.UUID) error {
	ch, err := rs.Conn.Channel()
	failOnError(err, "Failed to Open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rs.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body, err := json.Marshal(customer)
	failOnError(err, "Failed to json marshal Customer")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
			MessageId:   uuid.String(),
			Timestamp:   time.Now(),
			UserId:      "guest",
			AppId:       "sample",
		},
	)
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	return nil
}

*/
func failOnError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s: %s", msg, err)
	}
	return nil
}
