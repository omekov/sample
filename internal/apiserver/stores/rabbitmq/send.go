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

// // S
