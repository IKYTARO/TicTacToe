package application

import (
	"TicTacToe/internal/domain/model"
	"testing"
)

type GameOverTC struct {
	name  string
	board [model.BoardSize][model.BoardSize]model.CellState
}

func TestCheckGameOver_InProgress(t *testing.T) {
	tests := []GameOverTC{
		{
			name: "empty board",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "first move",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "mid game",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Cross, model.Empty},
				{model.Empty, model.Empty, model.Nought},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &model.Game{Board: model.Board{Cells: tt.board}}
			result := (&GameService{}).CheckGameOver(game)
			if result != model.InProgress {
				t.Errorf("expected InProgress, got %v", result)
			}
		})
	}
}

func TestCheckGameOver_CrossWins(t *testing.T) {
	tests := []GameOverTC{
		{
			name: "row 0",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Cross, model.Cross},
				{model.Nought, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "row 1",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Nought, model.Empty},
				{model.Cross, model.Cross, model.Cross},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "row 2",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Nought, model.Nought, model.Empty},
				{model.Cross, model.Cross, model.Cross},
			},
		},
		{
			name: "column 0",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Cross, model.Nought, model.Nought},
				{model.Cross, model.Empty, model.Empty},
			},
		},
		{
			name: "column 1",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Cross, model.Empty},
				{model.Empty, model.Cross, model.Empty},
				{model.Nought, model.Cross, model.Empty},
			},
		},
		{
			name: "column 2",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Nought, model.Cross},
				{model.Empty, model.Nought, model.Cross},
				{model.Empty, model.Empty, model.Cross},
			},
		},
		{
			name: "main diagonal",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Cross, model.Empty},
				{model.Nought, model.Empty, model.Cross},
			},
		},
		{
			name: "secondary diagonal",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Nought, model.Cross},
				{model.Empty, model.Cross, model.Nought},
				{model.Cross, model.Empty, model.Empty},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &model.Game{Board: model.Board{Cells: tt.board}}
			result := (&GameService{}).CheckGameOver(game)
			if result != model.CrossWon {
				t.Errorf("expected CrossWon, got %v", result)
			}
		})
	}
}

func TestCheckGameOver_NoughtWins(t *testing.T) {
	tests := []GameOverTC{
		{
			name: "row 0",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Nought, model.Nought},
				{model.Cross, model.Cross, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
		},
		{
			name: "column 1",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Empty},
				{model.Empty, model.Nought, model.Empty},
				{model.Cross, model.Nought, model.Empty},
			},
		},
		{
			name: "main diagonal",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Cross, model.Empty},
				{model.Empty, model.Nought, model.Cross},
				{model.Empty, model.Empty, model.Nought},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &model.Game{Board: model.Board{Cells: tt.board}}
			result := (&GameService{}).CheckGameOver(game)
			if result != model.NoughtWon {
				t.Errorf("expected NoughtWon, got %v", result)
			}
		})
	}
}

func TestCheckGameOver_Draw(t *testing.T) {
	board := [model.BoardSize][model.BoardSize]model.CellState{
		{model.Cross, model.Nought, model.Cross},
		{model.Cross, model.Nought, model.Nought},
		{model.Nought, model.Cross, model.Cross},
	}

	game := &model.Game{Board: model.Board{Cells: board}}
	result := (&GameService{}).CheckGameOver(game)
	if result != model.Draw {
		t.Errorf("expected Draw, got %v", result)
	}
}
