package application

import (
	"TicTacToe/internal/domain/model"
	"testing"

	"github.com/google/uuid"
)

type ValidationTC struct {
	name     string
	previous [model.BoardSize][model.BoardSize]model.CellState
	current  [model.BoardSize][model.BoardSize]model.CellState
}

func TestValidate_ValidMove(t *testing.T) {
	tests := []ValidationTC{
		{
			name: "first move",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "second move after computer",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Cross},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := uuid.New()
			previous := &model.Game{ID: id, Board: model.Board{Cells: tt.previous}}
			current := &model.Game{ID: id, Board: model.Board{Cells: tt.current}}

			err := (&GameService{}).Validate(previous, current)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidate_InvalidMoves(t *testing.T) {
	id := uuid.New()

	tests := []ValidationTC{
		{
			name: "no move made",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "two cells changed",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Cross, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "cell already occupied",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "put nought instead of cross",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "changed computer's previous move",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "invalid move count: two crosses added",
			previous: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			current: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Cross, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			previous := &model.Game{ID: id, Board: model.Board{Cells: tt.previous}}
			current := &model.Game{ID: id, Board: model.Board{Cells: tt.current}}

			err := (&GameService{}).Validate(previous, current)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}

func TestValidate_GameAlreadyFinished(t *testing.T) {
	id := uuid.New()

	previous := &model.Game{
		ID: id,
		Board: model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
			{model.Cross, model.Cross, model.Cross},
			{model.Nought, model.Nought, model.Empty},
			{model.Empty, model.Empty, model.Empty},
		}},
	}
	current := &model.Game{
		ID: id,
		Board: model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
			{model.Cross, model.Cross, model.Cross},
			{model.Nought, model.Nought, model.Empty},
			{model.Empty, model.Empty, model.Cross},
		}},
	}

	err := (&GameService{}).Validate(previous, current)
	if err == nil {
		t.Error("expected error for finished game, got nil")
	}
}

func TestValidate_DifferentIDs(t *testing.T) {
	previous := &model.Game{ID: uuid.New(), Board: model.Board{}}
	current := &model.Game{ID: uuid.New(), Board: model.Board{}}

	err := (&GameService{}).Validate(previous, current)
	if err == nil {
		t.Error("expected error for different IDs, got nil")
	}
}
