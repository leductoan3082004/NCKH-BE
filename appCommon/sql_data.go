package appCommon

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MgDBModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Status    int                `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

//func (sqlModel *SQLModel) GenUID(dbType DbType) {
//	uid := NewUID(uint32(sqlModel.Id), int(dbType), 271104)
//	sqlModel.FakeId = &uid
//}
//
//func (sqlModel *SQLModel) PrepareForInsert() {
//	now := time.Now().UTC()
//	sqlModel.Id = 0
//	sqlModel.CreatedAt = &now
//	sqlModel.UpdatedAt = &now
//}
//
//func (sqlModel *SQLModel) PrepareForUpdate() {
//	now := time.Now().UTC()
//	sqlModel.UpdatedAt = &now
//}
