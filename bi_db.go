package misbi

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"misbi/p"
	"time"
)

const (
	MONGO_HOST     = "127.0.0.1"
	MONGO_PORT     = "27017"
	MONGO_USER     = ""
	MONGO_PASSWORD = ""
	MONGO_DBNAME   = "userlog"
)

var mgoSession *mgo.Session

func InitDbConn() error {
	if mgoSession == nil {
		url := fmt.Sprintf("mongodb://%s:%s/%s", MONGO_HOST, MONGO_PORT, MONGO_DBNAME)
		session, err := mgo.DialWithTimeout(url, time.Duration(10)*time.Second)
		if err != nil {
			log.Println("[ERROR]Can not connect MongoDB, err:", err)
			return errors.New("Can not connect mongodb")
		}
		mgoSession = session
	}
	log.Println("Init mongo session succeed")
	return nil
}

func SaveBiLog(item *p.BiLog) {
	if mgoSession == nil {
		log.Println("MongoDB not connected, user log saved to file")
		log.Printf("[UserLog] %#v\n", item)
		return
	}
	db := mgoSession.DB(item.ProjectName)
	collection := db.C("userlog")
	item.Timestamp = time.Now().Unix()
	if err := collection.Insert(item); err != nil {
		log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
	}
}
