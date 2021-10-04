package dto

type CreateNewGameRequest struct {
	PlayerName      string `json:"playerName"`
	NumberOfMines   int    `json:"numberOfMines"`
	NumberOfColumns int    `json:"numberOfColumns"`
	NumberOfRows    int    `json:"numberOfRows"`
}

type ClickGameRequest struct {
	GameId string `json:"gameId"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
}
