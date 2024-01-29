package userstorage

import (
	"context"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

func (s *mongodbStore) Delete(ctx context.Context, conditions interface{}) error {
	db := s.db.Database(appCommon.MainDBName).Collection(postmodel.Post{}.TableName())
	_, err := db.DeleteOne(ctx, conditions)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
