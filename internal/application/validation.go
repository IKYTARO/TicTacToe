package application

import (
	domainerrors "TicTacToe/internal/domain/errors"
	"TicTacToe/internal/domain/model"
)

func (service *GameService) Validate(previous, current *model.Game) error {
	if err := service.validateGameInProgress(previous); err != nil {
		return err
	}

	row, col, err := validateSingleMove(previous, current)
	if err != nil {
		return err
	}

	if err = validateCellWasEmpty(previous, row, col); err != nil {
		return err
	}

	if err = validateMoveSymbol(current, row, col); err != nil {
		return err
	}

	if err = validatePreviousState(previous, current, row, col); err != nil {
		return err
	}

	if err = validateMoveCount(current); err != nil {
		return err
	}

	return nil
}

func (service *GameService) validateGameInProgress(game *model.Game) error {
	if service.CheckGameOver(game) != model.InProgress {
		return domainerrors.ErrGameAlreadyFinished
	}

	return nil
}

func validateSingleMove(previous, current *model.Game) (int, int, error) {
	var (
		found  bool
		rowIdx int
		colIdx int
	)

	for row := 0; row < model.BoardSize; row++ {
		for col := 0; col < model.BoardSize; col++ {
			if previous.Board.Cells[row][col] == current.Board.Cells[row][col] {
				continue
			}

			if found {
				return 0, 0, domainerrors.ErrInvalidBoard
			}

			found = true
			rowIdx = row
			colIdx = col
		}
	}

	if !found {
		return 0, 0, domainerrors.ErrInvalidMove
	}

	return rowIdx, colIdx, nil
}

func validateCellWasEmpty(previous *model.Game, row, col int) error {
	if previous.Board.Cells[row][col] != model.Empty {
		return domainerrors.ErrInvalidMove
	}

	return nil
}

func validateMoveSymbol(current *model.Game, row, col int) error {
	if current.Board.Cells[row][col] != model.Cross {
		return domainerrors.ErrInvalidMove
	}

	return nil
}

func validatePreviousState(previous, current *model.Game, changedRow, changedCol int) error {
	for row := 0; row < model.BoardSize; row++ {
		for col := 0; col < model.BoardSize; col++ {

			if row == changedRow && col == changedCol {
				continue
			}

			if previous.Board.Cells[row][col] != current.Board.Cells[row][col] {
				return domainerrors.ErrInvalidBoard
			}
		}
	}

	return nil
}

func validateMoveCount(game *model.Game) error {
	crosses := game.Board.Count(model.Cross)
	noughts := game.Board.Count(model.Nought)

	if crosses != noughts+1 {
		return domainerrors.ErrInvalidMove
	}

	return nil
}
