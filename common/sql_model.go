package common

import "go.mongodb.org/mongo-driver/bson/primitive"

type SqlModel struct {
	Status    int        `bson:"status,omitempty" json:"status"`
	CreatedAt *primitive.Timestamp `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt *primitive.Timestamp `bson:"updated_at,omitempty" json:"updated_at"`
}