package tt

import "go.mongodb.org/mongo-driver/v2/bson"

type BaseTemplate struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SysUserID int           `bson:"sys_user_id" json:"sys_user_id"`
	Display   int           `bson:"display" json:"display"`
	CreatedAt string        `bson:"created_at" json:"created_at"`
	UpdatedAt string        `bson:"updated_at" json:"updated_at"`
}

const (
	DisplayShow = 1
	DisplayHide = 0
)
