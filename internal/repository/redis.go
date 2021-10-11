package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"minesweeper-api/internal/model"
	"os"

	"github.com/gomodule/redigo/redis"
)

type RedisStore struct {
	redis.Conn
}

func NewRedisStore() RedisStore {
	c, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	return RedisStore{c}
}

func connect() (redis.Conn, error) {
	if url := os.Getenv("REDIS_CLOUD_URL"); url != "" {
		log.Print("[REDIS] - Connecting to cloud redis...")
		return redis.DialURL(url, redis.DialPassword(os.Getenv("REDIS_CLOUD_PASSWORD")))
	} else {
		log.Print("[REDIS] - Connecting to local redis...")
		return redis.DialURL("redis://localhost:6379")
	}
}

func (r RedisStore) SaveGame(g model.Game) error {
	key := fmt.Sprintf("game:%s", g.Id)
	gameJson, err := json.Marshal(g)
	if err != nil {
		return err
	}

	_, err = r.Do("SET", key, gameJson)

	if err != nil {
		return err
	}

	return nil
}

func (r RedisStore) SaveBoard(b model.Board) error {
	key := fmt.Sprintf("game:%s:board", b.GameId)
	gameJson, err := json.Marshal(b)
	if err != nil {
		return err
	}

	_, err = r.Do("SET", key, gameJson)

	if err != nil {
		return err
	}

	return nil
}

func (r RedisStore) FindGameById(id string) (*model.Game, error) {
	key := fmt.Sprintf("game:%s", id)
	value, err := redis.String(r.Do("GET", key))

	if err != nil {
		return nil, err
	}

	var game model.Game
	err = json.Unmarshal([]byte(value), &game)

	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (r RedisStore) FindBoardById(id string) (*model.Board, error) {
	key := fmt.Sprintf("game:%s:board", id)
	value, err := redis.String(r.Do("GET", key))

	if err != nil {
		return nil, err
	}

	var board model.Board
	err = json.Unmarshal([]byte(value), &board)

	if err != nil {
		return nil, err
	}

	return &board, nil
}

func (r RedisStore) DeleteBoardById(id string) {
	// TODO
}

func (r RedisStore) DeleteGameById(id string) {
	// TODO
}
