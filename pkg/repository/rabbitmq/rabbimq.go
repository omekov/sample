package rabbitmq

import (
	"time"

	"github.com/omekov/sample/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Client struct {
	logger              *logrus.Logger
	connConfig          *amqp.Config
	conn                *amqp.Connection
	channel             *amqp.Channel
	errorChannel        chan *amqp.Error
	config              config.RabbitMQ
	keepAlivePollPeriod int
}

func NewClient(keepAlivePollPeriod int, logger *logrus.Logger) *Client {
	return &Client{
		keepAlivePollPeriod: keepAlivePollPeriod,
		logger:              logger,
	}
}

func (rmq *Client) Close() {
	rmq.channel.Close()
	rmq.conn.Close()
	rmq.logger.Info("RMQ Closing connection")
}

func (rmq *Client) checkError(err error, message string) {
	if err == nil {
		return
	}

	rmq.logger.Errorf("%s:%v", message, err)
}

func (rmq *Client) GetConnection(url string) *amqp.Connection {
	attempt := 0
	ticker := time.NewTicker(time.Duration(rmq.keepAlivePollPeriod) * time.Second)
	defer ticker.Stop()
	for {
		if rmq.conn != nil {
			rmq.conn.Close()
		}

		rmq.logger.Printf("RMQ connecting on %s", url)
		conn, err := amqp.Dial(url)
		if err == nil {
			rmq.conn = conn
			rmq.errorChannel = make(chan *amqp.Error)
			rmq.conn.NotifyClose(rmq.errorChannel)
			attempt++
			rmq.logger.Info("RMQ connected success")

			rmq.openChannel()
			rmq.registerQueue()
			return rmq.conn
		}

		rmq.checkError(err, "RMQ connection failed")
		rmq.logger.Fatal("RMQ. Retrying connect in 3 sec...")
	}
}

func (rmq *Client) openChannel() {
	channel, err := rmq.conn.Channel()
	rmq.checkError(err, "RMQ Failed to open a channel.")
	rmq.channel = channel
}

func (rmq *Client) registerQueue() {
	exchangeType, exchangeDurable, exchangeAutoDelete := "direct", true, false
	exchangeName := "extest"
	rmq.logger.Printf("RMQ declaring exchange: exchange=%s, type=%s, durable=%t, autoDelete=%t", exchangeName, exchangeType, exchangeDurable, exchangeAutoDelete)
	err := rmq.channel.ExchangeDeclare(
		exchangeName,       // name
		exchangeType,       // type
		exchangeDurable,    // durable
		exchangeAutoDelete, // auto-deleted
		false,              // interval
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		rmq.checkError(err, "RMQ Failed to declare an exchange.")
	}

	queueDurable, queueExclusive, queueAutoDelete := true, false, false
	queueName := "qutest"
	rmq.logger.Printf("RMQ declaring queue=%s, durable=%t, exclusive=%t, autoDelete=%t", queueName, queueDurable, queueExclusive, queueAutoDelete)
	q, err := rmq.channel.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		false,
		nil,
	)
	rmq.config.QueueName = q.Name

	routingKey, noWait := "", false
	rmq.logger.Printf("RMQ binding queue: queue=%s, routingKey=%s, exchange=%s, noWait=%t", q.Name, routingKey, exchangeName, noWait)
	err = rmq.channel.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		noWait,
		nil,
	)
	if err != nil {
		rmq.checkError(err, "RMQ Failed to bind a queue.")
		return
	}

	if rmq.config.PrefetchCount > 0 {
		err = rmq.channel.Qos(rmq.config.PrefetchCount, 0, false)
		if err != nil {
			rmq.checkError(err, "RMQ Failed to set PrefetchCount.")
		}
		rmq.logger.Printf("RMQ Set PrefetchCount=%d", rmq.config.PrefetchCount)
	}
}
