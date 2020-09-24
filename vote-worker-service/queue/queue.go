package queue

import "github.com/mikestefanello/surveys-microservices/vote-service/vote"

// VoteQueue contains functions to receive votes from a queue
type VoteQueue interface {
	// Consume consumes votes from a queue and passes them through a
	// channel to be processed
	Consume(vc chan<- *vote.Vote)
}
