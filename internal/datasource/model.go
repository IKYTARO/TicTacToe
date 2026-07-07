package datasource

import (
	"TicTacToe/internal/domain/model"

	"github.com/google/uuid"
)

type GameEntity struct {
	ID    uuid.UUID
	Board [model.BoardSize][model.BoardSize]int
}
