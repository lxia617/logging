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
	MONGO_USER     = ""
	MONGO_PASSWORD = ""
	MONGO_DBNAME   = "userlog"
)

var MgoSession *mgo.Session

func InitDbConn(mongoHost, mongoPort string) error {
	url := fmt.Sprintf("mongodb://%s:%s/%s", mongoHost, mongoPort, MONGO_DBNAME)
	log.Println("Try to connect to MongoDB, url: ", url, "...")
	session, err := mgo.DialWithTimeout(url, time.Duration(10)*time.Second)
	if err != nil {
		log.Println("[ERROR]Can not connect MongoDB, err:", err)
		return errors.New("Can not connect mongodb")
	}
	MgoSession = session
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

func SaveDeviceInfo(item *p.DeviceInfo) error {
	if MgoSession == nil {
		log.Println("MongoDB not connected, user log saved to file")
		log.Printf("[UserLog] %#v\n", item)
		return errors.New("Can not connect mongoDb")
	}
	db := MgoSession.DB(item.NiVersion)
	collection := db.C("userlog")
	if err := collection.Insert(item); err != nil {
		log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
		return err
	}

	return nil
}
