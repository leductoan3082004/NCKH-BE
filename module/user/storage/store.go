package userstorage

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbStore struct {
	db *mongo.Client
}

func NewMgDBStore(db *mongo.Client) *mongodbStore {
	return &mongodbStore{
		db: db,
	}
}
