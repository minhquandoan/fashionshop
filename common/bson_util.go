package common

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func ConvertToBsonD(i *interface{}) (*bson.D, error) {

	if *i == nil {
		log.Println("Nil here")
		return &bson.D{{"status", 1}}, nil
	}

	b, err := bson.Marshal(i)
    if err != nil {
        return nil, err
    }
    var d bson.D
    err = bson.Unmarshal(b, &d)
    if err != nil {
        return nil, err
    }
    return &d, nil
}
