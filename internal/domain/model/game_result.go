package model

type GameResult int

const (
	InProgress GameResult = 0
	CrossWon   GameResult = 1
	NoughtWon  GameResult = 2
	Draw       GameResult = 3
)
