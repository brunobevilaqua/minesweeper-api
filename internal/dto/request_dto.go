package dto

type CreateNewGameRequest struct {
	PlayerName      string `json:"playerName"`
	NumberOfMines   int    `json:"numberOfMines"`
	NumberOfColumns int    `json:"numberOfColumns"`
	NumberOfRows    int    `json:"numberOfRows"`
}

type ActionRequest struct {
	Action string `json:"action"` // flag or click
	Row    int    `json:"row"`
	Column int    `json:"column"`
}
