package imagemodel

import "nckh-BE/appCommon"

const EntityName = "Image"

type Image struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	Url                 string `json:"url" bson:"url"`
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
