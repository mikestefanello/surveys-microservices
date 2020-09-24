package storage

import "github.com/mikestefanello/surveys-microservices/vote-service/vote"

type VoteStorage interface {
	Insert(v *vote.Vote) error

	UpdateResults(v *vote.Vote) error
}
