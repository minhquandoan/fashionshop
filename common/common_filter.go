package common

import "go.mongodb.org/mongo-driver/bson/primitive"

type FilterStatus struct {
	Status int `bson:"status,omitempty" json:"status"`
}

type FilterCreatedTime struct {
	CreatedAt primitive.Timestamp `bson:"created_at,omitempty" json:"created_at"`
}

type FilterUpdatedTime struct {
	UpdatedAt primitive.Timestamp `bson:"updated_at,omitempty" json:"updated_at"`
}
