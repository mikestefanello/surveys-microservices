package repository

import (
	"fmt"

	"github.com/mikestefanello/surveys-microservices/vote-service/config"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/streadway/amqp"
)

type rabbitVoteRepository struct {
	cfg        config.RabbitConfig
	serializer vote.Serializer
}

// NewRabbitVoteWriterRepository creates a new RabbitMQ vote writer repository
func NewRabbitVoteWriterRepository(cfg config.RabbitConfig, sz vote.Serializer) (vote.WriterRepository, error) {
	r := &rabbitVoteRepository{
		cfg:        cfg,
		serializer: sz,
	}
	conn, ch, err := r.connect()
	if err != nil {
		return r, err
	}
	defer conn.Close()
	defer ch.Close()
	return r, nil
}

func (r *rabbitVoteRepository) connect() (*amqp.Connection, *amqp.Channel, error) {
	// Connect to rabbit
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/", r.cfg.Username, r.cfg.Password, r.cfg.Hostname, r.cfg.Port)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	// Specify the queue
	_, err = ch.QueueDeclare(
		r.cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

func (r *rabbitVoteRepository) Insert(v *vote.Vote) error {
	// Encode the vote
	enc, err := r.serializer.Encode(v)
	if err != nil {
		return err
	}

	// Connect to rabbit
	conn, ch, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	// Publish the vote
	err = ch.Publish(
		"",
		r.cfg.QueueName,
		true,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  r.serializer.GetContentType(),
			Body:         enc,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
