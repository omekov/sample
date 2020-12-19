package rabbitmq

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/streadway/amqp"
// )

// type RabbitMQ struct {
// 	connection   *amqp.Connection
// 	channel      *amqp.Channel
// 	errorChannel chan *amqp.Error
// 	config       Config
// }

// // type Config struct {
// // 	URL           string
// // 	ExchangeName  string
// // 	QueueName     string
// // 	PrefetchCount int
// // }

// func NewRabbitMQ() {
// 	amqp.Config
// }

// func (rmq *RabbitMQ) Close() {
// 	rmq.channel.Close()
// 	rmq.connection.Close()
// }
// func (rmq *RabbitMQ) checkError(err error, message string) {
// 	if err == nil {
// 		return
// 	}
// 	fmt.Errorf("%s:%v", message, err)
// }

// func (rmq *RabbitMQ) connect() {
// 	for {
// 		if rmq.connection != nil {
// 			rmq.connection.Close()
// 		}
// 		url := "amqp://guest:guest@localhost:5672/test"
// 		log.Printf("RMQ connecting on %s", url)
// 		conn, err := amqp.Dial(url)
// 		if err == nil {
// 			rmq.connection = conn
// 			rmq.errorChannel = make(chan *amqp.Error)
// 			rmq.connection.NotifyClose(rmq.errorChannel)

// 			log.Println("RMQ connected success")

// 			rmq.openChannel()
// 			rmq.registerQueue()
// 			return
// 		}
// 		rmq.checkError(err, "RMQ connection failed")
// 		log.Fatal("RMQ. Retrying connect in 3 sec...")
// 		time.Sleep(3000 * time.Millisecond)
// 	}
// }

// func (rmq *RabbitMQ) openChannel() {
// 	channel, err := rmq.connection.Channel()
// 	rmq.checkError(err, "RMQ Failed to open a channel.")
// 	rmq.channel = channel
// }

// func (rmq *RabbitMQ) registerQueue() {
// 	exchangeType, exchangeDurable, exchangeAutoDelete := "direct", true, false
// 	exchangeName := "extest"
// 	log.Printf("RMQ declaring exchange: exchange=%s, type=%s, durable=%t, autoDelete=%t", exchangeName, exchangeType, exchangeDurable, exchangeAutoDelete)
// 	err := rmq.channel.ExchangeDeclare(
// 		exchangeName,       // name
// 		exchangeType,       // type
// 		exchangeDurable,    // durable
// 		exchangeAutoDelete, // auto-deleted
// 		false,              // interval
// 		false,              // no-wait
// 		nil,                // arguments
// 	)
// 	if err != nil {
// 		rmq.checkError(err, "RMQ Failed to declare an exchange.")
// 	}
// 	queueDurable, queueExclusive, queueAutoDelete := true, false, false
// 	queueName := "qutest"
// 	log.Printf("RMQ declaring queue=%s, durable=%t, exclusive=%t, autoDelete=%t", queueName, queueDurable, queueExclusive, queueAutoDelete)
// 	q, err := rmq.channel.QueueDeclare(
// 		queueName,
// 		queueDurable,
// 		queueAutoDelete,
// 		queueExclusive,
// 		false,
// 		nil,
// 	)
// 	rmq.config.QueueName = q.Name

// 	routingKey, noWait := "", false
// 	log.Printf("RMQ binding queue: queue=%s, routingKey=%s, exchange=%s, noWait=%t", q.Name, routingKey, exchangeName, noWait)
// 	err = rmq.channel.QueueBind(
// 		q.Name,
// 		routingKey,
// 		exchangeName,
// 		noWait,
// 		nil,
// 	)
// 	if err != nil {
// 		rmq.checkError(err, "RMQ Failed to bind a queue.")
// 		return
// 	}
// 	if rmq.config.PrefetchCount > 0 {
// 		err = rmq.channel.Qos(rmq.config.PrefetchCount, 0, false)
// 		if err != nil {
// 			rmq.checkError(err, "RMQ Failed to set PrefetchCount.")
// 		}
// 		log.Printf("RMQ Set PrefetchCount=%d", rmq.config.PrefetchCount)
// 	}
// }
