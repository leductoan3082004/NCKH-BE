package feedbackstorage

import (
	"context"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

func (s *mongodbStore) Delete(ctx context.Context, conditions interface{}) error {
	db := s.db.Database(appCommon.MainDBName).Collection(feedbackmodel.Feedback{}.TableName())
	_, err := db.DeleteOne(ctx, conditions)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
