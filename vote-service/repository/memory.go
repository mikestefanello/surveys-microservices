package repository

import (
	"sync"

	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
)

type voteMemoryRepository struct {
	storage map[string]*vote.Vote
	mutex   *sync.Mutex
}

// NewMemoryVoteWriterRepository creates a new vote writer repository that stores in memory
func NewMemoryVoteWriterRepository() (vote.WriterRepository, error) {
	return &voteMemoryRepository{
		storage: make(map[string]*vote.Vote),
		mutex:   &sync.Mutex{},
	}, nil
}

func (r *voteMemoryRepository) Insert(v *vote.Vote) error {
	r.mutex.Lock()
	r.storage[v.ID] = v
	r.mutex.Unlock()
	return nil
}
