package datasource

import (
	"TicTacToe/internal/domain/model"
	"testing"

	"github.com/google/uuid"
)

func TestStorage_StoreAndLoad(t *testing.T) {
	storage := NewStorage()
	id := uuid.New()
	entity := &GameEntity{
		ID:    id,
		Board: [model.BoardSize][model.BoardSize]int{},
	}

	storage.Store(id, entity)

	loaded, ok := storage.Load(id)
	if !ok {
		t.Fatal("expected to find stored entity")
	}

	if loaded.ID != entity.ID {
		t.Errorf("ID mismatch: expected %v, got %v", entity.ID, loaded.ID)
	}
}

func TestStorage_LoadNotFound(t *testing.T) {
	storage := NewStorage()
	id := uuid.New()

	_, ok := storage.Load(id)
	if ok {
		t.Error("expected false for non-existent key")
	}
}

func TestStorage_Overwrite(t *testing.T) {
	storage := NewStorage()
	id := uuid.New()

	first := &GameEntity{ID: id, Board: [model.BoardSize][model.BoardSize]int{}}
	second := &GameEntity{
		ID:    id,
		Board: [model.BoardSize][model.BoardSize]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}

	storage.Store(id, first)
	storage.Store(id, second)

	loaded, ok := storage.Load(id)
	if !ok {
		t.Fatal("expected to find stored entity")
	}

	if loaded.Board[0][0] != 1 {
		t.Errorf("expected updated board, got %v", loaded.Board)
	}
}
