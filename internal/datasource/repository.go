package datasource

import (
	domainerrors "TicTacToe/internal/domain/errors"
	"TicTacToe/internal/domain/model"

	"github.com/google/uuid"
)

type GameRepository struct {
	storage *Storage
}

func NewGameRepository(storage *Storage) *GameRepository {
	return &GameRepository{
		storage: storage,
	}
}

func (r *GameRepository) Save(game *model.Game) {
	entity := DomainToEntity(game)
	r.storage.Store(entity.ID, entity)
}

func (r *GameRepository) FindByID(id uuid.UUID) (*model.Game, error) {
	entity, ok := r.storage.Load(id)
	if !ok {
		return nil, domainerrors.ErrGameNotFound
	}
	return EntityToDomain(entity), nil
}
