package storage

import "github.com/mikestefanello/surveys-microservices/vote-service/vote"

// VoteStorage contains functions to store votes and the results of votes
type VoteStorage interface {
	// Insert inserts a vote in to storage
	Insert(v *vote.Vote) error

	// UpdateResults updates the total vote count results for the survey and
	// question of a given vote
	UpdateResults(v *vote.Vote) error
}
