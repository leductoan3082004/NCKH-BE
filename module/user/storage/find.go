package userstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	usermodel "nckh-BE/module/user/model"
)

func (s *mongodbStore) Find(ctx context.Context, conditions interface{}, fields ...string) (*usermodel.User, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(usermodel.User{}.TableName())
	var data usermodel.User
	err := db.FindOne(ctx, conditions).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
