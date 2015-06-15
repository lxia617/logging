package be

import (
	"errors"
	"fmt"
	"github.com/MISingularity/logging/p"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	MONGO_HOST     = "127.0.0.1"
	MONGO_PORT     = "27017"
	MONGO_USER     = ""
	MONGO_PASSWORD = ""
	MONGO_DBNAME   = "userlog"
)

var MgoSession *mgo.Session

func InitDbConn() error {
	if MgoSession == nil {
		url := fmt.Sprintf("mongodb://%s:%s/%s", MONGO_HOST, MONGO_PORT, MONGO_DBNAME)
		session, err := mgo.DialWithTimeout(url, time.Duration(10)*time.Second)
		if err != nil {
			log.Println("[ERROR]Can not connect MongoDB, err:", err)
			return errors.New("Can not connect mongodb")
		}
		MgoSession = session
	}
	log.Println("Init mongo session succeed")
	return nil
}

func SaveBiLog(item *p.BiLog) error {
	if MgoSession == nil {
		log.Println("MongoDB not connected, user log saved to file")
		log.Printf("[UserLog] %#v\n", item)
		return errors.New("Can not connect mongoDb")
	}
	db := MgoSession.DB(item.ProjectName)
	collection := db.C("userlog")
	if err := collection.Insert(item); err != nil {
		log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
		return err
	}

	return nil
}
