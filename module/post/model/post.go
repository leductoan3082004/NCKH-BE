package postmodel

import "nckh-BE/appCommon"

const EntityName = "Post"

type Post struct {
	appCommon.MgDBModel `json:",inline"`
	Title               string `json:"title" bson:"title"`
	Content             string `json:"content" bson:"content"`
	Author              string `json:"author" bson:"author"`
}

type PostCreate struct {
	Title   string `json:"title" binding:"title"`
	Content string `json:"content" binding:"content"`
	Author  string `json:"author" binding:"author"`
}
type PostUpdate struct {
	PostId  string  `json:"post_id" binding:"required"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Author  *string `json:"author"`
}

func (Post) TableName() string {
	return "post"
}
