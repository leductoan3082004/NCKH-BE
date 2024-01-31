package imagebiz

import (
	"context"
	"fmt"
	"github.com/lequocbinh04/go-sdk/logger"
	"github.com/lequocbinh04/go-sdk/plugin/aws"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nckh-BE/appCommon"
	imagemodel "nckh-BE/module/image/model"
	"path/filepath"
	"time"
)

type uploadStore interface {
	Create(ctx context.Context, data *imagemodel.Image) error
}
type uploadBiz struct {
	store  uploadStore
	logger logger.Logger
	s3     aws.S3
}

func NewUploadBiz(store uploadStore, s3 aws.S3) *uploadBiz {
	return &uploadBiz{
		store:  store,
		s3:     s3,
		logger: logger.GetCurrent().GetLogger("ImageUploadBiz"),
	}
}

func (biz *uploadBiz) UploadImage(ctx context.Context, data []byte, path, fileName string) (*imagemodel.Image, error) {
	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		return nil, imagemodel.ErrInvalidImageFormat
	}
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	_, err := biz.s3.UploadFileData(ctx, data, appCommon.Join("/", path, fileName))
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, imagemodel.ErrCannotUploadImage(err)
	}

	res := &imagemodel.Image{
		MgDBModel: appCommon.MgDBModel{
			Id:        primitive.ObjectID{},
			Status:    0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Url: appCommon.Join("/", appCommon.S3Domain, path, fileName),
	}
	if err := biz.store.Create(ctx, res); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(imagemodel.EntityName, err)
	}
	return res, nil
}
