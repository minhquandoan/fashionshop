package common

type Image struct {
	Url       string `bson:"url,omitempty"`
	Width     int `bson:"width,omitempty"`
	Height    int `bson:"height,omitempty"`
	CloudName string `bson:"cloudname,omitempty"`
	Extension string `bson:"extension,omitempty"`
}