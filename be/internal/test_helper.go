package internal

import (
	"bytes"
	"github.com/MiSingularity/logging/be"
	"github.com/MiSingularity/logging/p"
	"gopkg.in/mgo.v2/bson"
	"log"
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
		be.InitDbConn()
	}
	db := be.MgoSession.DB(item.ProjectName)
	collection := db.C("userlog")
	info, err := collection.RemoveAll(bson.M{"projectname": item.ProjectName, "actionname": item.ActionName, "timestamp": item.Timestamp, "detail": item.Detail})
	log.Printf("remove result: %#v, %#v \n", info, err)
}
