package userstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
)

func (s *mongodbStore) GetUser(ctx context.Context, userId string) (*usermodel.User, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	db := s.db.Database(appCommon.MainDBName).Collection(usermodel.User{}.TableName())
	var data usermodel.User

	conditions := bson.M{
		"_id": id,
	}

	err = db.FindOne(ctx, conditions).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
