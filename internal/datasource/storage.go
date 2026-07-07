package datasource

import (
	"sync"

	"github.com/google/uuid"
)

type Storage struct {
	data sync.Map
}

func NewStorage() *Storage {
	return &Storage{
		data: sync.Map{},
	}
}

func (s *Storage) Store(id uuid.UUID, game *GameEntity) {
	s.data.Store(id, game)
}

func (s *Storage) Load(id uuid.UUID) (*GameEntity, bool) {
	value, ok := s.data.Load(id)
	if !ok {
		return nil, false
	}

	entity, ok := value.(*GameEntity)
	if !ok {
		return nil, false
	}

	return entity, true
}
