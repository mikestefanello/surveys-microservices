package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/config"
)

type postgresVoteStorage struct {
	config     config.PostgresConfig
	connection *pgx.Conn
}

// NewPostgresVoteStorage creates a new Postgres vote storage
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
	q := fmt.Sprintf("INSERT INTO %s(id, survey, question, created) VALUES($1, $2, $3, $4)", p.config.Tables.Votes)
	_, err := p.connection.Exec(context.Background(), q, v.ID, v.Survey, v.Question, v.Timestamp)
	return err
}

func (p *postgresVoteStorage) UpdateResults(v *vote.Vote) error {
	// Check if results already exist for the vote's survey and question
	var r int
	q := fmt.Sprintf("SELECT votes FROM %s WHERE survey = $1 AND question = $2", p.config.Tables.Results)
	err := p.connection.QueryRow(context.Background(), q, v.Survey, v.Question).Scan(&r)

	switch err {
	case nil:
		// Increment the results
		q = fmt.Sprintf("UPDATE %s SET votes = votes + 1, last_update = $1", p.config.Tables.Results)
		_, err := p.connection.Exec(context.Background(), q, time.Now().UTC().Unix())
		if err != nil {
			return err
		}
	case pgx.ErrNoRows:
		// Initialize the results
		q = fmt.Sprintf("INSERT INTO %s(survey, question, votes, last_update) VALUES($1, $2, $3, $4)", p.config.Tables.Results)
		_, err := p.connection.Exec(context.Background(), q, v.Survey, v.Question, 1, time.Now().UTC().Unix())
		if err != nil {
			return err
		}
	default:
		return err
	}

	return nil
}
