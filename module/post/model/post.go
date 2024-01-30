package postmodel

import "nckh-BE/appCommon"

const EntityName = "Post"

type Post struct {
	appCommon.MgDBModel `json:",inline"`
	Title               string   `json:"title" bson:"title"`
	Content             string   `json:"content" bson:"content"`
	Author              string   `json:"author" bson:"author"`
	ImageUrl            []string `json:"image_url" bson:"image_url"`
	Type                string   `json:"type" bson:"type"`
}

type PostCreate struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Type    string `json:"type" binding:"required"`
}
type PostUpdate struct {
	PostId  string  `json:"post_id" binding:"required"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Author  *string `json:"author"`
	Type    *string `json:"type"`
}
type PostDelete struct {
	PostId string `json:"post_id" binding:"required"`
}

func (Post) TableName() string {
	return "post"
}
