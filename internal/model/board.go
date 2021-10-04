package model

const (
	MAX_DEFAULT_ROWS  = 30
	MAX_DEFAULT_COLS  = 30
	MIN_DEFAULT_ROWS  = 8
	MIN_DEFAULT_COLS  = 8
	MIN_DEFAULT_MINES = 12
	MAX_DEFAULT_MINES = 12
)

type Board struct {
	Rows  int        `json:"rows"`
	Cols  int        `json:"columns"`
	Mines int        `json:"mines"`
	Grid  []CellGrid `json:"grid"`
}

type CellGrid []Cell

type Cell struct {
	Mine, Clicked bool
	Value         int
}

func NewBoard(rows, cols, mines int) *Board {
	b := Board{}

	b.Cols = setValue(cols, MAX_DEFAULT_COLS, MIN_DEFAULT_COLS)
	b.Rows = setValue(cols, MAX_DEFAULT_COLS, MIN_DEFAULT_ROWS)
	b.Mines = setValue(cols, MAX_DEFAULT_MINES, MIN_DEFAULT_MINES)

	return &b
}

func setValue(v, max, min int) int {
	// if input value is greater than MAX_DEFAULT
	if v > max {
		return max
	}

	// if input value is less than MIN_DEFAULT
	if v < min {
		return min
	}

	return v
}
