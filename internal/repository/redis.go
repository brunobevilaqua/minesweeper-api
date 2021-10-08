package repository

import (
	"encoding/json"
	"log"
	"minesweeper-api/internal/model"
	"os"

	"github.com/gomodule/redigo/redis"
)

type RedisStore struct {
	redis.Conn
}

func (r RedisStore) Save(g model.Game) (*model.Game, error) {
	gameJson, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	_, err = r.Do("SET", g.Id, gameJson)

	if err != nil {
		return nil, err
	}

	return r.FindById(g.Id)
}

func (r RedisStore) FindById(id string) (*model.Game, error) {
	value, err := redis.String(r.Do("GET", id))

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

func (r RedisStore) FindByPlayerName(name string) (*model.Game, error) {
	value, err := redis.String(r.Do("GET", "player.name:"+name))

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
