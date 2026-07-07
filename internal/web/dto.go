package web

type GameRequest struct {
	Board [][]int `json:"board"`
}

type GameResponse struct {
	ID     string  `json:"id"`
	Board  [][]int `json:"board"`
	Status string  `json:"status"`
}

const (
	StatusInProgress = "in_progress"
	StatusCrossWon   = "cross_won"
	StatusNoughtWon  = "nought_won"
	StatusDraw       = "draw"
)
