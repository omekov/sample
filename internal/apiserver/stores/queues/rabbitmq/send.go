package rabbitmq

// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/omekov/sample/internal/apiserver/models"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/streadway/amqp"
// )

// // type Config struct {
// // 	URL           string
// // }

// type RabbitServer struct {
// 	conn          amqp.Connection
// 	exchangeName  string
// 	queueName     string
// 	prefetchCount int
// }

// func failOnError(err error, msg string) error {
// 	if err != nil {
// 		return fmt.Errorf("%s: %s", msg, err)
// 	}
// 	return nil
// }
// func (rs *RabbitServer) Connect(config Config) error {
// 	conn, err := amqp.Dial(config.URL)
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()
// 	rs.conn = *conn
// 	rs.exchangeName = ""
// 	rs.queueName = "SAMPLE_CUSTOMER_ACTIVED"
// 	return nil
// }

// // Send ...
// func (rs *RabbitServer) Send(customer *models.Customer, uuid uuid.UUID) error {
// 	ch, err := rs.conn.Channel()
// 	failOnError(err, "Failed to Open a channel")
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		rs.queueName,
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	body, err := json.Marshal(customer)
// 	failOnError(err, "Failed to json marshal Customer")

// 	err = ch.Publish(
// 		"",
// 		q.Name,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        []byte(body),
// 			MessageId:   uuid.String(),
// 			Timestamp:   time.Now(),
// 			UserId:      "guest",
// 			AppId:       "sample",
// 		},
// 	)
// 	log.Printf(" [x] Sent %s", body)
// 	failOnError(err, "Failed to publish a message")
// 	return nil
// }
