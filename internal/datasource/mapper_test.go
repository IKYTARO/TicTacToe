package datasource

import (
	"TicTacToe/internal/domain/model"
	"testing"

	"github.com/google/uuid"
)

func TestMapper_DomainToEntity(t *testing.T) {
	id := uuid.New()
	domainGame := &model.Game{
		ID: id,
		Board: model.Board{
			Cells: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Cross, model.Empty},
				{model.Empty, model.Empty, model.Nought},
			},
		},
	}

	entity := DomainToEntity(domainGame)

	if entity.ID != id {
		t.Errorf("ID mismatch: expected %v, got %v", id, entity.ID)
	}

	if entity.Board[0][0] != int(model.Cross) {
		t.Errorf("expected Cross at (0,0), got %d", entity.Board[0][0])
	}
	if entity.Board[0][1] != int(model.Nought) {
		t.Errorf("expected Nought at (0,1), got %d", entity.Board[0][1])
	}
	if entity.Board[0][2] != int(model.Empty) {
		t.Errorf("expected Empty at (0,2), got %d", entity.Board[0][2])
	}
}

func TestMapper_EntityToDomain(t *testing.T) {
	id := uuid.New()
	entity := &GameEntity{
		ID: id,
		Board: [model.BoardSize][model.BoardSize]int{
			{1, 2, 0},
			{0, 1, 2},
			{2, 0, 1},
		},
	}

	domainGame := EntityToDomain(entity)

	if domainGame.ID != id {
		t.Errorf("ID mismatch: expected %v, got %v", id, domainGame.ID)
	}

	if domainGame.Board.Cells[0][0] != model.Cross {
		t.Errorf("expected Cross at (0,0), got %v", domainGame.Board.Cells[0][0])
	}
	if domainGame.Board.Cells[0][1] != model.Nought {
		t.Errorf("expected Nought at (0,1), got %v", domainGame.Board.Cells[0][1])
	}
	if domainGame.Board.Cells[0][2] != model.Empty {
		t.Errorf("expected Empty at (0,2), got %v", domainGame.Board.Cells[0][2])
	}
}

func TestMapper_RoundTrip(t *testing.T) {
	id := uuid.New()
	original := &model.Game{
		ID: id,
		Board: model.Board{
			Cells: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Cross, model.Nought},
				{model.Nought, model.Empty, model.Cross},
			},
		},
	}

	entity := DomainToEntity(original)
	restored := EntityToDomain(entity)

	if restored.ID != original.ID {
		t.Errorf("ID mismatch after round trip")
	}

	for row := 0; row < model.BoardSize; row++ {
		for col := 0; col < model.BoardSize; col++ {
			if restored.Board.Cells[row][col] != original.Board.Cells[row][col] {
				t.Errorf("mismatch at (%d, %d): expected %v, got %v",
					row, col, original.Board.Cells[row][col], restored.Board.Cells[row][col])
			}
		}
	}
}
