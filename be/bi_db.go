package be

import (
	"errors"
	"fmt"
	"github.com/MISingularity/logging/p"
	"gopkg.in/mgo.v2"
	"log"
	"time"
	"encoding/json"
	"strings"
)

const (
	MONGO_USER     = ""
	MONGO_PASSWORD = ""
	MONGO_DBNAME   = "userlog"
)

type GoBiLog struct {
	ProjectName string
	ActionName  string
	Timestamp   int64
	Detail string
}

type Action struct {
	Command string `json:"command"`
	Actions string `json:"actions"`
}

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
	log.Printf("[UserLog] %#v\n", item)
	if MgoSession == nil {
		log.Println("MongoDB not connected, user log saved to file")
		log.Printf("[UserLog] %#v\n", item)
		return errors.New("Can not connect mongoDb")
	}
	db := MgoSession.DB(item.ProjectName)
	collection := db.C(item.ActionName)

	if item.ActionName ==  "tracking_data"{
		act := Action{}
		detail := string(item.Detail[:])
		detail = strings.Replace(detail,"'","\"",-1)
		log.Println("[Act from]:")
		log.Println(detail)
		json.Unmarshal([]byte(detail), &act)
		log.Println("[Act Result]:")
		log.Println(act)
		log.Println(act.Command)
		log.Println("[Act.actions]" + act.Actions)

		if err := collection.Insert(act); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	}

/*	goBiItem := &GoBiLog{
		ProjectName:item.ProjectName,
		ActionName:item.ActionName,
		Timestamp:item.Timestamp,
		Detail:string(item.Detail[:]),
	}

	log.Printf("[GoBiLog All] %#v\n", goBiItem.Detail)

	log.Printf("[GoBiLog All] %#v\n", goBiItem)

	if err := collection.Insert(goBiItem); err != nil {
		log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
		return err
	}*/

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
