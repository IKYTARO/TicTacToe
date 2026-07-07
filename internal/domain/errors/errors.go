package domainerrors

import "errors"

var (
	ErrInvalidMove         = errors.New("invalid move")
	ErrNoAvailableMoves    = errors.New("no available moves")
	ErrGameNotFound        = errors.New("game not found")
	ErrGameAlreadyFinished = errors.New("game already finished")
	ErrInvalidBoard        = errors.New("invalid board")
)
