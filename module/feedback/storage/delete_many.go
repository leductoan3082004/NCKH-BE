package feedbackstorage

import (
	"context"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

func (s *mongodbStore) DeleteMany(ctx context.Context, conditions interface{}) error {
	db := s.db.Database(appCommon.MainDBName).Collection(feedbackmodel.Feedback{}.TableName())
	_, err := db.DeleteMany(ctx, conditions)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
