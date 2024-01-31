package poststorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

func (s *mongodbStore) Find(ctx context.Context, conditions interface{}, fields ...string) (*postmodel.Post, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(postmodel.Post{}.TableName())
	var data postmodel.Post
	err := db.FindOne(ctx, conditions).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
