package application

import (
	domainerrors "TicTacToe/internal/domain/errors"
	"TicTacToe/internal/domain/model"
	"math"
)

const (
	winScore  = 10
	drawScore = 0
	loseScore = -10
)

// Move описывает возможный ход и его оценку.
type Move struct {
	Row int
	Col int
}

func (service *GameService) NextMove(game *model.Game) error {
	move, err := bestMove(&game.Board)
	if err != nil {
		return err
	}

	game.Board.Cells[move.Row][move.Col] = model.Nought

	return nil
}

func bestMove(board *model.Board) (Move, error) {
	bestScore := math.MinInt
	bestMove := Move{
		Row: -1,
		Col: -1,
	}

	for _, move := range availableMoves(board) {
		board.Cells[move.Row][move.Col] = model.Nought
		score := minimax(board, 0, false)
		board.Cells[move.Row][move.Col] = model.Empty

		if score > bestScore {
			bestScore = score
			bestMove.Row = move.Row
			bestMove.Col = move.Col
		}
	}

	if bestMove.Row == -1 {
		return Move{}, domainerrors.ErrNoAvailableMoves
	}
	return bestMove, nil
}

func minimax(board *model.Board, depth int, isMaximizing bool) int {
	if score, terminal := evaluate(board); terminal {
		switch {
		case score > 0:
			return score - depth
		case score < 0:
			return score + depth
		default:
			return drawScore
		}
	}

	if isMaximizing {
		best := math.MinInt
		for _, move := range availableMoves(board) {
			board.Cells[move.Row][move.Col] = model.Nought
			score := minimax(board, depth+1, false)
			board.Cells[move.Row][move.Col] = model.Empty
			if score > best {
				best = score
			}
		}
		return best
	}

	best := math.MaxInt
	for _, move := range availableMoves(board) {
		board.Cells[move.Row][move.Col] = model.Cross
		score := minimax(board, depth+1, true)
		board.Cells[move.Row][move.Col] = model.Empty
		if score < best {
			best = score
		}
	}
	return best
}

func evaluate(board *model.Board) (score int, terminal bool) {
	winner := findWinner(board)

	if winner == model.Nought {
		return winScore, true
	}

	if winner == model.Cross {
		return loseScore, true
	}

	if board.Count(model.Empty) == 0 {
		return drawScore, true
	}

	return 0, false
}

func availableMoves(board *model.Board) []Move {
	var moves []Move
	for row := 0; row < model.BoardSize; row++ {
		for col := 0; col < model.BoardSize; col++ {
			if board.Cells[row][col] == model.Empty {
				moves = append(moves, Move{
					Row: row,
					Col: col})
			}
		}
	}
	return moves
}
