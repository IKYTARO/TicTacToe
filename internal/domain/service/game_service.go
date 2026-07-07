package service

import "TicTacToe/internal/domain/model"

type GameService interface {
	// CreateGame : Метод создания новой игры
	CreateGame() *model.Game

	// ProcessGame : Метод обработки текущего состояния игры
	ProcessGame(current *model.Game) (model.GameResult, error)

	// NextMove : Метод получения следующего хода текущей игры алгоритмом «Минимакс»
	NextMove(game *model.Game) error

	// Validate : Метод валидации игрового поля текущей игры
	Validate(previous, current *model.Game) error

	// CheckGameOver : Метод проверки окончания игры
	CheckGameOver(game *model.Game) model.GameResult
}
