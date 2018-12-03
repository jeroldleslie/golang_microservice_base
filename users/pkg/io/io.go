package io

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username"  bson:"username"`
	Password  string        `json:"password"  bson:"password"`
	Firstname string        `json:"firstname"  bson:"firstname"`
	Lastname  string        `json:"lastname"  bson:"lastname"`
	Usertype  string        `json:"usertype"  bson:"usertype"`
	Email     string        `json:"email"  bson:"email"`
	Phone     string        `json:"phone"  bson:"phone"`
}

func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}


type Authentication struct {
	Username  string        `json:"username"  bson:"username"`
	Password  string        `json:"password"  bson:"password"`
}

func (a Authentication) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}