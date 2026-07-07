package model

type Board struct {
	Cells [BoardSize][BoardSize]CellState
}

func (board *Board) Count(state CellState) int {
	count := 0
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if board.Cells[row][col] == state {
				count++
			}
		}
	}
	return count
}
