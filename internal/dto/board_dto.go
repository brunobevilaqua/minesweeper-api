package dto

type BoardDto struct {
	Columns       int `json:"columns"`
	Rows          int `json:"rows"`
	NumberOfMines int `json:"numberOfMines"`
	Clicks        int `json:"clicks"`
}

func NewBoardDto(row, col, numberOfMines, clicks int) BoardDto {
	return BoardDto{
		Columns:       col,
		Rows:          row,
		NumberOfMines: numberOfMines,
		Clicks:        clicks,
	}
}
