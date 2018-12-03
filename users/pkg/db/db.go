package db

import (
	"gopkg.in/mgo.v2"
	"github.com/go-kit/kit/log"
	)

var mgoSession *mgo.Session
var mgoUsersCollectionSession *mgo.Session
var mongo_conn_str = "mongodb:27017"

// Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongo_conn_str)
		mgoSession.SetMode(mgo.Monotonic, true)
		if err != nil {
			return nil, err
		}
	}
	return mgoSession.Clone(), nil
}

func EnsureIndexOnAllCollections(logger log.Logger) error {
	// TODO find a better way to ensure mondodb index in application init
	session, err := GetMongoSession()
	if err != nil {
		logger.Log("err",err)
		return err
	}
	c := session.DB("fivekilometer").C("users")
	index := &mgo.Index{
		Key:        []string{"id", "username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	logger.Log("EnsureIndex","Ensuring index for users collection")
	err = c.EnsureIndex(*index)
	if err != nil {
		logger.Log("err",err)
		return err
	}
	return nil
}

func GetUsersCollection(logger log.Logger)(c *mgo.Collection,err error) {
	session, err := GetMongoSession()
	if err != nil {
		logger.Log("err",err)
		return nil, err
	}
	return session.DB("fivekilometer").C("users"),err
}





