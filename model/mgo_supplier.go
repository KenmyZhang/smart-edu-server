package model

import (
	"gopkg.in/mgo.v2"
)

var QuestionCollection *mgo.Collection
var mgoSession *mgo.Session
func NewMgoSupplier(mgoEndpoint string) *mgo.Session {
	var err error
	mgoSession, err = mgo.Dial(mgoEndpoint)
	if err != nil {
		panic(err)
	}

	// Optional Switch the MgoSession to a monotonic behavior.
	mgoSession.SetMode(mgo.Monotonic, true)
	InitMgoCollections()
	return mgoSession
}

func InitMgoCollections() {
	QuestionCollection = mgoSession.DB("smart_edu").C("questions")
	/*	QuestionCollection.EnsureIndex(mgo.Index{
			Key:    []string{"drugNum"},
			Unique: true,
		})*/
}

