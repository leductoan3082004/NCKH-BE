package poststorage

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/mapstructure"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
	"time"
)

func (s *mongodbStore) Update(ctx context.Context, conditions interface{}, data interface{}) error {
	db := s.db.Database(appCommon.MainDBName).Collection(postmodel.Post{}.TableName())

	var updateData bson.M
	if err := mapstructure.Decode(data, &updateData); err != nil {
		return appCommon.ErrInternal(err)
	}
	if updateData["$set"] != nil {
		updateData["$set"].(bson.M)["updated_at"] = time.Now()
	} else {
		updateData["$set"] = bson.M{
			"updated_at": time.Now(),
		}
	}
	data = updateData

	_, err := db.UpdateOne(ctx, conditions, data)
	if err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
