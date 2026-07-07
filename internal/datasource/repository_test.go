package datasource

import (
	"TicTacToe/internal/domain/model"
	"testing"

	"github.com/google/uuid"
)

func TestRepository_SaveAndFindByID(t *testing.T) {
	storage := NewStorage()
	repo := NewGameRepository(storage)

	id := uuid.New()
	game := &model.Game{
		ID: id,
		Board: model.Board{
			Cells: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
	}

	repo.Save(game)

	found, err := repo.FindByID(id)
	if err != nil {
		t.Fatalf("expected to find game, got error: %v", err)
	}

	if found.ID != id {
		t.Errorf("ID mismatch: expected %v, got %v", id, found.ID)
	}

	if found.Board.Cells[0][0] != model.Cross {
		t.Errorf("expected Cross at (0,0), got %v", found.Board.Cells[0][0])
	}
	if found.Board.Cells[0][1] != model.Nought {
		t.Errorf("expected Nought at (0,1), got %v", found.Board.Cells[0][1])
	}
}

func TestRepository_FindByID_NotFound(t *testing.T) {
	storage := NewStorage()
	repo := NewGameRepository(storage)

	id := uuid.New()

	_, err := repo.FindByID(id)
	if err == nil {
		t.Error("expected error for non-existent game, got nil")
	}
}

func TestRepository_SaveUpdatesExisting(t *testing.T) {
	storage := NewStorage()
	repo := NewGameRepository(storage)

	id := uuid.New()

	game1 := &model.Game{
		ID:    id,
		Board: model.Board{},
	}
	repo.Save(game1)

	game2 := &model.Game{
		ID: id,
		Board: model.Board{
			Cells: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
	}
	repo.Save(game2)

	found, err := repo.FindByID(id)
	if err != nil {
		t.Fatalf("expected to find game, got error: %v", err)
	}

	if found.Board.Cells[0][0] != model.Cross {
		t.Errorf("expected Cross at (0,0) after update, got %v", found.Board.Cells[0][0])
	}
	if found.Board.Cells[1][1] != model.Nought {
		t.Errorf("expected Nought at (1,1) after update, got %v", found.Board.Cells[1][1])
	}
}

func TestRepository_SaveReturnsCopy(t *testing.T) {
	storage := NewStorage()
	repo := NewGameRepository(storage)

	id := uuid.New()
	game := &model.Game{
		ID: id,
		Board: model.Board{
			Cells: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
	}

	repo.Save(game)

	game.Board.Cells[0][0] = model.Nought

	found, err := repo.FindByID(id)
	if err != nil {
		t.Fatalf("expected to find game, got error: %v", err)
	}

	if found.Board.Cells[0][0] != model.Cross {
		t.Errorf("storage was affected by external modification: expected Cross, got %v",
			found.Board.Cells[0][0])
	}
}
