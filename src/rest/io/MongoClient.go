package io

import (
	"gopkg.in/mgo.v2"
)

type MongoClient struct {
	Session *mgo.Session
	Db      string
}

func NewMongoClient() (mc *MongoClient) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	mc = &MongoClient{Session: session, Db: "rest"}
	return
}

func (mc *MongoClient) Close() {
	mc.Session.Close()
}

func (mc *MongoClient) Insert(collection string, data interface{}) error {
	//data must be an address
	return mc.Session.DB(mc.Db).C(collection).Insert(data)
}
