package io

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Password  string        `json:"password,omitempty"  bson:"password,omitempty"`
	Firstname string        `json:"firstname,omitempty"  bson:"firstname,omitempty"`
	Lastname  string        `json:"lastname,omitempty"  bson:"lastname,omitempty"`
	Usertype  string        `json:"usertype,omitempty"  bson:"usertype,omitempty"`
	Email     string        `json:"email,omitempty"  bson:"email,omitempty"`
	Phone     string        `json:"phone,omitempty"  bson:"phone,omitempty"`
}

func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}

type Authentication struct {
	Email    string `json:"email"  bson:"email"`
	Password string `json:"password"  bson:"password"`
}

func (a Authentication) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
