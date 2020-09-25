package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/mikestefanello/surveys-microservices/vote-service/config"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
)

type postgresResultsRepository struct {
	config     config.PostgresConfig
	connection *pgx.Conn
}

// NewPostgresResultsRepository creates a new Postgres vote results repository
func NewPostgresResultsRepository(cfg config.PostgresConfig) (vote.ResultsRepository, error) {
	p := &postgresResultsRepository{config: cfg}

	err := p.connect()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *postgresResultsRepository) connect() error {
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

func (p *postgresResultsRepository) GetResults(surveyID string) (vote.Results, error) {
	// Query to get the results for this survey.
	q := fmt.Sprintf("SELECT * FROM %s WHERE survey = $1", p.config.Tables.Results)
	rows, _ := p.connection.Query(context.Background(), q, surveyID)
	results := vote.Results{
		Survey: surveyID,
	}

	for rows.Next() {
		var question int
		var votes int
		var lastUpdate int64

		// Extract the data from the row
		err := rows.Scan(nil, &question, &votes, &lastUpdate)
		if err != nil {
			return results, err
		}

		// Add the results for this question
		result := vote.Result{
			Question: question,
			Votes:    votes,
		}
		results.Results = append(results.Results, result)

		// Update the last update time
		if results.UpdatedAt < lastUpdate {
			results.UpdatedAt = lastUpdate
		}
	}

	// Check if there are no votes, which could also mean the survey ID is invalid
	if results.UpdatedAt == 0 {
		return results, vote.ErrResultsNotFound
	}

	return results, rows.Err()
}
