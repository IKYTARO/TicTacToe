package datasource

import "TicTacToe/internal/domain/model"

func DomainToEntity(domainGame *model.Game) *GameEntity {
	var entityBoard [model.BoardSize][model.BoardSize]int
	for row := 0; row < model.BoardSize; row++ {
		for col := 0; col < model.BoardSize; col++ {
			entityBoard[row][col] = int(domainGame.Board.Cells[row][col])
		}
	}

	return &GameEntity{
		ID:    domainGame.ID,
		Board: entityBoard,
	}
}

func EntityToDomain(entityGame *GameEntity) *model.Game {
	var gameBoard model.Board
	for row := 0; row < model.BoardSize; row++ {
		for col := 0; col < model.BoardSize; col++ {
			gameBoard.Cells[row][col] = model.CellState(entityGame.Board[row][col])
		}
	}

	return &model.Game{
		ID:    entityGame.ID,
		Board: gameBoard,
	}
}
