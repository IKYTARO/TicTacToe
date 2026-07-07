package application

import (
	"TicTacToe/internal/domain/model"
	"TicTacToe/internal/domain/repository"
)

type GameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (service *GameService) CreateGame() *model.Game {
	newGame := model.NewGame()
	service.repo.Save(newGame)
	return newGame
}

func (service *GameService) ProcessGame(current *model.Game) (model.GameResult, error) {
	previous, err := service.repo.FindByID(current.ID)
	if err != nil {
		return 0, err
	}

	if err = service.Validate(previous, current); err != nil {
		return 0, err
	}

	service.repo.Save(current)

	if result := service.CheckGameOver(current); result != model.InProgress {
		return result, nil
	}

	if err = service.NextMove(current); err != nil {
		return 0, err
	}

	service.repo.Save(current)

	return service.CheckGameOver(current), nil
}
