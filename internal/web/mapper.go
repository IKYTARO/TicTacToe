package web

import (
	"TicTacToe/internal/domain/model"
	"fmt"

	"github.com/google/uuid"
)

// RequestToDomain : преобразует запрос клиента в доменную модель.
func RequestToDomain(id string, request *GameRequest) (*model.Game, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	if len(request.Board) != model.BoardSize {
		return nil, fmt.Errorf("invalid board size: %d", len(request.Board))
	}

	var cells [model.BoardSize][model.BoardSize]model.CellState

	for row := 0; row < model.BoardSize; row++ {
		if len(request.Board[row]) != model.BoardSize {
			return nil, fmt.Errorf("invalid board size: %d", len(request.Board[row]))
		}
		for col := 0; col < model.BoardSize; col++ {
			cells[row][col] = model.CellState(request.Board[row][col])
		}
	}

	return &model.Game{
		ID:    uid,
		Board: model.Board{Cells: cells},
	}, nil
}

// DomainToResponse : преобразует доменную модель в ответ клиенту.
func DomainToResponse(game *model.Game, result model.GameResult) *GameResponse {
	board := make([][]int, model.BoardSize)
	for row := 0; row < model.BoardSize; row++ {
		board[row] = make([]int, model.BoardSize)
		for col := 0; col < model.BoardSize; col++ {
			board[row][col] = int(game.Board.Cells[row][col])
		}
	}

	return &GameResponse{
		ID:     game.ID.String(),
		Board:  board,
		Status: gameResultToString(result),
	}
}

func gameResultToString(result model.GameResult) string {
	switch result {
	case model.InProgress:
		return StatusInProgress
	case model.CrossWon:
		return StatusCrossWon
	case model.NoughtWon:
		return StatusNoughtWon
	case model.Draw:
		return StatusDraw
	default:
		return StatusInProgress
	}
}
