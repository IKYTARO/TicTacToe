package application

import (
	"TicTacToe/internal/domain/model"
	"testing"
)

type MinimaxTC struct {
	name     string
	board    [model.BoardSize][model.BoardSize]model.CellState
	score    int
	terminal bool
}

func TestEvaluate(t *testing.T) {
	tests := []MinimaxTC{
		{
			name: "nought wins",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Nought, model.Nought, model.Nought},
				{model.Cross, model.Cross, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			score:    winScore,
			terminal: true,
		},
		{
			name: "cross wins",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Cross, model.Cross},
				{model.Nought, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			score:    loseScore,
			terminal: true,
		},
		{
			name: "draw",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Nought, model.Cross},
				{model.Cross, model.Nought, model.Nought},
				{model.Nought, model.Cross, model.Cross},
			},
			score:    drawScore,
			terminal: true,
		},
		{
			name: "in progress",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Cross, model.Empty, model.Empty},
				{model.Empty, model.Nought, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			score:    0,
			terminal: false,
		},
		{
			name: "empty board",
			board: [model.BoardSize][model.BoardSize]model.CellState{
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
				{model.Empty, model.Empty, model.Empty},
			},
			score:    0,
			terminal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &model.Board{Cells: tt.board}
			score, terminal := evaluate(board)

			if terminal != tt.terminal {
				t.Errorf("terminal: expected %v, got %v", tt.terminal, terminal)
			}
			if score != tt.score {
				t.Errorf("score: expected %d, got %d", tt.score, score)
			}
		})
	}
}

func TestBestMove_BlocksOpponentWin(t *testing.T) {
	board := &model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
		{model.Cross, model.Cross, model.Empty},
		{model.Nought, model.Empty, model.Empty},
		{model.Empty, model.Empty, model.Empty},
	}}

	move, err := bestMove(board)
	if err != nil {
		t.Fatalf("expected move, got error: %v", err)
	}

	expectedRow, expectedCol := 0, 2
	if move.Row != expectedRow || move.Col != expectedCol {
		t.Errorf("expected (%d, %d), got (%d, %d)", expectedRow, expectedCol, move.Row, move.Col)
	}
}

func TestBestMove_TakesWin(t *testing.T) {
	board := &model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
		{model.Nought, model.Nought, model.Empty},
		{model.Cross, model.Cross, model.Empty},
		{model.Empty, model.Empty, model.Empty},
	}}

	move, err := bestMove(board)
	if err != nil {
		t.Fatalf("expected move, got error: %v", err)
	}

	expectedRow, expectedCol := 0, 2
	if move.Row != expectedRow || move.Col != expectedCol {
		t.Errorf("expected (%d, %d), got (%d, %d)", expectedRow, expectedCol, move.Row, move.Col)
	}
}

func TestBestMove_CenterPreference(t *testing.T) {
	board := &model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
		{model.Empty, model.Empty, model.Empty},
		{model.Empty, model.Empty, model.Empty},
		{model.Empty, model.Empty, model.Empty},
	}}

	move, err := bestMove(board)
	if err != nil {
		t.Fatalf("expected move, got error: %v", err)
	}

	if move.Row < 0 || move.Row >= model.BoardSize || move.Col < 0 || move.Col >= model.BoardSize {
		t.Errorf("invalid move: (%d, %d)", move.Row, move.Col)
	}
}

func TestBestMove_NoAvailableMoves(t *testing.T) {
	board := &model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
		{model.Cross, model.Nought, model.Cross},
		{model.Cross, model.Nought, model.Nought},
		{model.Nought, model.Cross, model.Cross},
	}}

	_, err := bestMove(board)
	if err == nil {
		t.Error("expected error for full board, got nil")
	}
}

func TestMinimax_EmptyBoard(t *testing.T) {
	board := &model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
		{model.Empty, model.Empty, model.Empty},
		{model.Empty, model.Empty, model.Empty},
		{model.Empty, model.Empty, model.Empty},
	}}

	score := minimax(board, 0, true)
	if score < loseScore || score > winScore {
		t.Errorf("score out of range: %d", score)
	}

	if score != 0 {
		t.Errorf("expected draw score (0) for empty board, got %d", score)
	}
}

func TestMinimax_TerminalPosition(t *testing.T) {
	board := &model.Board{Cells: [model.BoardSize][model.BoardSize]model.CellState{
		{model.Nought, model.Nought, model.Nought},
		{model.Cross, model.Cross, model.Empty},
		{model.Empty, model.Empty, model.Empty},
	}}

	score := minimax(board, 0, true)
	if score != winScore {
		t.Errorf("expected winScore (%d), got %d", winScore, score)
	}
}
