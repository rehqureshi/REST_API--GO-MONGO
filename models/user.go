package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

func IsValidObjectId(id string) bool {
	return bson.IsObjectIdHex(id)
}

func ToObjectId(id string) bson.ObjectId {
	return bson.ObjectIdHex(id)
}

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}
