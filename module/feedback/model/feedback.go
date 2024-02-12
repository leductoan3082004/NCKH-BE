package feedbackmodel

import "nckh-BE/appCommon"

const EntityName = "Feedback"

type Feedback struct {
	appCommon.MgDBModel `json:",inline" bson:",inline"`
	Content             string `json:"content" bson:"content"`
	Topic               string `json:"topic" bson:"topic"`
	Email               string `json:"email" bson:"email"`
	Phone               string `json:"phone" bson:"phone"`
	Name                string `json:"name" bson:"name"`
}

func (Feedback) TableName() string {
	return "feedback"
}

type FeedbackCreate struct {
	Content string `json:"content" binding:"required"`
	Topic   string `json:"topic" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"numeric,min=9,max=20"`
	Name    string `json:"name" binding:"required,min=1,max=100"`
}

type FeedbackDelete struct {
	FeedbackIds []string `json:"feedback_ids" binding:"required"`
}
