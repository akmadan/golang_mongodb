package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId
	Name   string
	Gender string
	Age    int
}
