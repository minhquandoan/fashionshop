package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paging struct {
	Page int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
	Total int `json:"total" form:"total"`

	//Support cursor with UID
	FakeCursor primitive.ObjectID `json:"cursor" form:"cursor"`
	NextCursor primitive.ObjectID `json:"next_cursor"`
}

func (p *Paging) Fullfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}