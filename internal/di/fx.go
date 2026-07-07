package di

import (
	"TicTacToe/internal/application"
	"TicTacToe/internal/datasource"
	"TicTacToe/internal/domain/repository"
	"TicTacToe/internal/domain/service"
	"TicTacToe/internal/web"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(datasource.NewStorage),

	fx.Provide(fx.Annotate(
		datasource.NewGameRepository,
		fx.As(new(repository.GameRepository)),
	),
	),

	fx.Provide(fx.Annotate(
		application.NewGameService,
		fx.As(new(service.GameService)),
	),
	),

	fx.Provide(web.NewGameHandler),
)
