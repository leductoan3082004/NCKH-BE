package imagemodel

import (
	"nckh-BE/appCommon"
	"time"
)

const EntityName = "Image"

type Image struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	Url                 string `json:"url" bson:"url"`
}
type ImageDelete struct {
	ImageId string `json:"image_id" binding:"required"`
}
type ImageList struct {
	TimeFrom *int64 `form:"time_from"`
	TimeTo   *int64 `form:"time_to"`
}

func (s *ImageList) FulFill() {
	if s.TimeTo == nil {
		s.TimeTo = new(int64)
		*s.TimeTo = time.Now().Unix()
	}
	if s.TimeFrom == nil {
		s.TimeFrom = new(int64)
		*s.TimeFrom = 0
	}
}

func (Image) TableName() string {
	return "image"
}

var (
	ErrCannotUploadImage = func(err error) error {
		return appCommon.NewCustomError(500, err, "Cannot upload image", "ErrCannotUploadImage")
	}
	ErrOverLimitUploadImage = appCommon.NewCustomError(400, nil, "Over limit upload image", "ErrOverLimitUploadImage")
	ErrInvalidImageFormat   = appCommon.NewCustomError(400, nil, "Invalid image format", "ErrInvalidImageFormat")
)
