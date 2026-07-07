package repository

import (
	"TicTacToe/internal/domain/model"

	"github.com/google/uuid"
)

type GameRepository interface {
	// Save : Метод сохранения текущей игры
	Save(game *model.Game)

	// FindByID : Метод получения текущей игры
	FindByID(id uuid.UUID) (*model.Game, error)
}
