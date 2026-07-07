package application

import "TicTacToe/internal/domain/model"

func (service *GameService) CheckGameOver(game *model.Game) model.GameResult {
	board := &game.Board

	if winner := findWinner(board); winner != model.Empty {
		return toGameResult(winner)
	}

	if board.Count(model.Empty) == 0 {
		return model.Draw
	}

	return model.InProgress
}

func findWinner(board *model.Board) model.CellState {
	if winner := checkRows(board); winner != model.Empty {
		return winner
	}

	if winner := checkColumns(board); winner != model.Empty {
		return winner
	}

	return checkDiagonals(board)
}

func checkRows(board *model.Board) model.CellState {
	for row := 0; row < model.BoardSize; row++ {
		if winner := checkRow(board, row); winner != model.Empty {
			return winner
		}
	}

	return model.Empty
}

func checkColumns(board *model.Board) model.CellState {
	for col := 0; col < model.BoardSize; col++ {
		if winner := checkColumn(board, col); winner != model.Empty {
			return winner
		}
	}

	return model.Empty
}

func checkDiagonals(board *model.Board) model.CellState {
	if winner := checkMainDiagonal(board); winner != model.Empty {
		return winner
	}

	return checkSecondaryDiagonal(board)
}

func checkRow(board *model.Board, row int) model.CellState {
	first := board.Cells[row][0]

	if first == model.Empty {
		return model.Empty
	}

	for col := 1; col < model.BoardSize; col++ {
		if board.Cells[row][col] != first {
			return model.Empty
		}
	}

	return first
}

func checkColumn(board *model.Board, col int) model.CellState {
	first := board.Cells[0][col]

	if first == model.Empty {
		return model.Empty
	}

	for row := 1; row < model.BoardSize; row++ {
		if board.Cells[row][col] != first {
			return model.Empty
		}
	}

	return first
}

func checkMainDiagonal(board *model.Board) model.CellState {
	first := board.Cells[0][0]

	if first == model.Empty {
		return model.Empty
	}

	for i := 1; i < model.BoardSize; i++ {
		if board.Cells[i][i] != first {
			return model.Empty
		}
	}

	return first
}

func checkSecondaryDiagonal(board *model.Board) model.CellState {
	first := board.Cells[0][model.BoardSize-1]

	if first == model.Empty {
		return model.Empty
	}

	for i := 1; i < model.BoardSize; i++ {
		if board.Cells[i][model.BoardSize-1-i] != first {
			return model.Empty
		}
	}

	return first
}

func toGameResult(state model.CellState) model.GameResult {
	switch state {
	case model.Cross:
		return model.CrossWon

	case model.Nought:
		return model.NoughtWon

	default:
		return model.InProgress
	}
}
