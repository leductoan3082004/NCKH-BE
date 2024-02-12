package feedbackstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
)

func (s *mongodbStore) Find(ctx context.Context, conditions interface{}, fields ...string) (*feedbackmodel.Feedback, error) {
	db := s.db.Database(appCommon.MainDBName).Collection(feedbackmodel.Feedback{}.TableName())
	var data feedbackmodel.Feedback
	err := db.FindOne(ctx, conditions).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
