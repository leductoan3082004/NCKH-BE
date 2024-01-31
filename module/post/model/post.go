package postmodel

import "nckh-BE/appCommon"

const EntityName = "Post"

type MetaData struct {
	Title    string   `json:"title" bson:"title"`
	ImageUrl string   `json:"image_url" bson:"image_url"`
	Tag      []string `json:"tag" bson:"tag"`
	Category []string `json:"category" bson:"category"`
}

type Post struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	MetaData            `json:",inline" bson:",inline"`
	Content             string `json:"content" bson:"content"`
	Author              string `json:"author" bson:"author"`
}

type SimplePost struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	MetaData            `json:",inline" bson:",inline"`
}

type PostCreate struct {
	ImageUrl string   `json:"image_url"`
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Author   string   `json:"author" binding:"required"`
	Tag      []string `json:"tag" bson:"tag"`
	Category []string `json:"category" bson:"category"`
}
type PostUpdate struct {
	PostId   string    `json:"post_id" binding:"required"`
	Title    *string   `json:"title"`
	Content  *string   `json:"content"`
	Author   *string   `json:"author"`
	ImageUrl *string   `json:"image_url"`
	Tag      *[]string `json:"tag" bson:"tag"`
	Category *[]string `json:"category" bson:"category"`
}
type PostDelete struct {
	PostId string `json:"post_id" binding:"required"`
}
type PostList struct {
	Tag      *string `json:"tag" form:"tag"`
	Category *string `json:"category" form:"category"`
	Content  *string `json:"content" form:"content"`
}
type PostFind struct {
	PostId string `json:"post_id" binding:"required"`
}

func (Post) TableName() string {
	return "post"
}
