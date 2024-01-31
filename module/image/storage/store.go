package imagestorage

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type sqlStore struct {
	db  *mongo.Client
	rdb *redis.Client
}

func NewMgDBStore(db *mongo.Client) *sqlStore {
	return &sqlStore{
		db:  db,
		rdb: nil,
	}
}

func NewSQLStoreWithRedis(db *mongo.Client, rdb *redis.Client) *sqlStore {
	return &sqlStore{
		db:  db,
		rdb: rdb,
	}
}
