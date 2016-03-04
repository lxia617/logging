package testhelper

import (
	"bytes"
	"log"

	"github.com/MISingularity/logging/be"
	"github.com/MISingularity/logging/p"
	"gopkg.in/mgo.v2/bson"
)

func IsEqual(a, b *p.BiLog) bool {
	if a == nil || b == nil {
		return false
	}
	if a.ProjectName != b.ProjectName {
		return false
	}
	if a.ActionName != b.ActionName {
		return false
	}
	if a.Timestamp != b.Timestamp {
		return false
	}
	if !bytes.Equal(a.Detail, b.Detail) {
		return false
	}
	return true
}

func RetrieveBiLog(item *p.BiLog) *p.BiLog {
	db := be.MgoSession.DB(item.ProjectName)
	collection := db.C("userlog")

	var ret p.BiLog
	err := collection.Find(bson.M{"projectname": item.ProjectName, "actionname": item.ActionName, "timestamp": item.Timestamp, "detail": item.Detail}).One(&ret)
	if err != nil {
		return nil
	}

	log.Printf("retrieved bi log: %#v\n", ret)
	return &ret
}

func DeleteBiLog(item *p.BiLog) {
	if be.MgoSession == nil {
		be.InitDbConn("127.0.0.1", "27017")
	}
	db := be.MgoSession.DB(item.ProjectName)
	collection := db.C("userlog")
	info, err := collection.RemoveAll(bson.M{"projectname": item.ProjectName, "actionname": item.ActionName, "timestamp": item.Timestamp, "detail": item.Detail})
	log.Printf("remove result: %#v, %#v \n", info, err)
}
