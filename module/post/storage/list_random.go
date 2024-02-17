package poststorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"nckh-BE/appCommon"
	postmodel "nckh-BE/module/post/model"
)

func (s *mongodbStore) ListRandom(ctx context.Context, condition bson.M, paging *appCommon.Paging, moreInfo ...string) ([]postmodel.SimplePost, error) {
	collection := s.db.Database(appCommon.MainDBName).Collection(postmodel.Post{}.TableName())

	if paging == nil {
		paging = &appCommon.Paging{
			Page:       1,
			FakeCursor: "",
			Limit:      50,
		}
	}
	pipeline := []bson.M{
		{"$match": condition},
		{"$sample": bson.M{"size": paging.Limit}},
	}
	opts := options.Aggregate().SetBatchSize(int32(paging.Limit))

	cursor, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		log.Fatal(err)
	}

	var res []postmodel.SimplePost
	if err = cursor.All(ctx, &res); err != nil {
		return nil, appCommon.ErrDB(err)
	}

	if res == nil {
		return []postmodel.SimplePost{}, nil
	}
	return res, nil
}
