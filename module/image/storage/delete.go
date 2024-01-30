package imagestorage

import (
	"context"
	"nckh-BE/appCommon"
	imagemodel "nckh-BE/module/image/model"
)

func (s *sqlStore) Delete(ctx context.Context, conditions interface{}) error {
	db := s.db.Database(appCommon.MainDBName).Collection(imagemodel.Image{}.TableName())
	_, err := db.DeleteOne(ctx, conditions)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
