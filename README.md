# Minesweeper-Api
Simple minesweeper REST Api developed using Go Lang and also:
- Docker
- Redis
- Gin

## Run
1. ```make up``` or ```docker-compose up```
2. ```make run```
3. Import `minesweeper-api.postman_collection.json` file on Postman

## Endpoints:
There are endpoints exposed to create new game, get game and to update game by sending click or flag actions. 
***All services responses are the same*** – Check `api_response_sample.json` in project root for a response example.

### 1. Create Game 
Create a new game by passing `player name`, `row`, `cols` and `number of mines`. In the service side there is a check to validate if the values are valid comparing with some default values for row, col and mines.

**Default Values:**
```
MAX_DEFAULT_ROWS  = 30
MAX_DEFAULT_COLS  = 30
MIN_DEFAULT_ROWS  = 5
MIN_DEFAULT_COLS  = 5
MIN_DEFAULT_MINES = 32
MAX_DEFAULT_MINES = 8
```
- Endpoint: `POST /api/games`
- cURL: ```curl --location --request POST 'http://localhost:8080/api/games/' \
--header 'Content-Type: application/json' \
--data-raw '{
	"playerName": "Bruno Bevilaqua",
	"numberOfMines": 8,
	"numberOfColumns": 5,
	"numberOfRows": 5
}' ```


### 2. Get Game By Game Id
- Endpoint: `GET /api/games/:id`
- cURL: ```curl --location --request GET 'http://localhost:8080/api/games/eb61a4eb-fb16-41e3-b923-ab95391934e2'```

### 3. Send a Click of Flag Action
Request should have `row`, `col` and also the `action`. Possible actions are: `"action": "click"` or `"action": "flag"`.
Sending a request with `"action":"flag"` for a cell that is already flagged will then remove the flag of it. 

- Endpoint: `PUT /api/games/:id`
- cURL: ```curl --location --request PUT 'http://localhost:8080/api/games/eb61a4eb-fb16-41e3-b923-ab95391934e2' \
--header 'Content-Type: application/json' \
--data-raw '{
	"action": "flag", 
	"row": 3,
	"column": 0
}'```
