package imagestorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	imagemodel "nckh-BE/module/image/model"
	"time"
)

func (s *sqlStore) Create(ctx context.Context, data *imagemodel.Image) error {
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
