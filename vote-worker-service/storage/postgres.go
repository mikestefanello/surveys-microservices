package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/config"
)

type postgresVoteStorage struct {
	config     config.PostgresConfig
	connection *pgx.Conn
}

func NewPostgresVoteStorage(cfg config.PostgresConfig) (VoteStorage, error) {
	p := &postgresVoteStorage{config: cfg}

	err := p.connect()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *postgresVoteStorage) connect() error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		p.config.User,
		p.config.Password,
		p.config.Hostname,
		p.config.Port,
		p.config.Database,
	)
	conn, err := pgx.Connect(context.Background(), addr)
	if err != nil {
		return err
	}

	p.connection = conn
	return nil
}

func (p *postgresVoteStorage) Insert(v *vote.Vote) error {
	q := fmt.Sprintf("INSERT INTO %s(id, survey, question, created) values($1, $2, $3, $4)", p.config.Tables.Votes)
	_, err := p.connection.Exec(context.Background(), q, v.ID, v.Survey, v.Question, v.Timestamp)
	return err
}

func (p *postgresVoteStorage) UpdateResults(v *vote.Vote) error {
	return nil
}
