package poststorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	feedbackmodel "nckh-BE/module/feedback/model"
	"time"
)

func (s *mongodbStore) Create(ctx context.Context, data *feedbackmodel.Feedback) error {
	db := s.db.Database(appCommon.MainDBName).Collection(data.TableName())
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	newData, err := db.InsertOne(ctx, data)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	data.Id = newData.InsertedID.(primitive.ObjectID)
	return nil
}
