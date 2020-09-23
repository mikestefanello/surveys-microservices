package repository

import (
	"fmt"

	"github.com/mikestefanello/surveys-microservices/vote-service/config"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/streadway/amqp"
)

type rabbitVoteRepository struct {
	cfg config.RabbitConfig
}

func NewRabbitVoteWriterRepository(cfg config.RabbitConfig) (vote.WriterRepository, error) {
	return &rabbitVoteRepository{cfg: cfg}, nil
}

func (r *rabbitVoteRepository) connect() (*amqp.Connection, *amqp.Channel, error) {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/", r.cfg.Username, r.cfg.Password, r.cfg.Hostname, r.cfg.Port)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

func (r *rabbitVoteRepository) Insert(v *vote.Vote) error {
	conn, ch, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	_, err = ch.QueueDeclare(
		r.cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		r.cfg.QueueName,
		true,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			// TODO: Send real data
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
