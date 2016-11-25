package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type MOAddr struct {
  session, table, doc string
} 

type MOValue struct {
  k, v string
  key, value string
}

type MOGValue struct {
	Key, Value string
}

func mongo_insert(moaddress MOAddr, movalue MOValue) bool{
  session, err := mgo.Dial(moaddress.session)
  if err != nil {
    return false
  }
  defer session.Close()
  c := session.DB(moaddress.table).C(moaddress.doc)
  c.Insert(bson.M{movalue.k: movalue.key, movalue.v: movalue.value})
  return true
}

func mongo_export(moaddress MOAddr, movalue MOValue) []MOGValue {
	session, err := mgo.Dial(moaddress.session)
	if err != nil {
			panic(err)
	}
	defer session.Close()
	c := session.DB(moaddress.table).C(moaddress.doc)
	result := []MOGValue{}
	iter := c.Find(bson.M{movalue.k: movalue.key}).Limit(100).Iter()
	err = iter.All(&result)
	if err != nil {
			log.Fatal(err)
	}
	return result
}