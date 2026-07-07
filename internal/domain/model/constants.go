package model

type CellState int

const (
	Empty  CellState = 0
	Cross  CellState = 1
	Nought CellState = 2
)

const (
	BoardSize = 3
)
