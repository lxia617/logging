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

type ServiceStartInfo struct {
	DeviceId string `json:"deviceId"`
	NiVersion string `json:"niVersion"`
	Model string `json:"model"`
	Manufacture string `json:"manufacture"`
}

type ServiceRemoveInfo struct {
	DeviceId string `json:"deviceId"`
	NiVersion string `json:"niVersion"`
}

type SearchPerformanceInfo struct {
	DeviceId string `json:"deviceId"`
	NiVersion string `json:"niVersion"`
	QueryTimestamp int64 `json:"queryTimestamp"`
	PerformPathCount int32 `json:"performPathCount"`
	ActionListCount int32 `json:"actionListCount"`
}

type NiPerformanceInfo struct {
	DeviceId string `json:"deviceId"`
	NiVersion string `json:"niVersion"`
	QueryTimestamp int64 `json:"queryTimestamp"`
	QueryVoiceTimestamp int64 `json:"queryVoiceTimestamp"`
	ReceiveVoiceResultTimestamp int64 `json:"receiveVoiceResultTimestamp"`
	SendQueryTimestamp int64 `json:"sendQueryTimestamp"`
	ReceiveQueryResultTimestamp int64 `json:"receiveQueryResultTimestamp"`
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

	detail := string(item.Detail[:])
	detail = strings.Replace(detail,"'","\"",-1)

	if item.ActionName ==  "service_start_info"{
		service_start_info := ServiceStartInfo{}
		json.Unmarshal([]byte(detail), &service_start_info)
		log.Println("[Result]:")
		log.Println(service_start_info)

		if err := collection.Insert(service_start_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	}else if item.ActionName == "service_remove_info"{
		service_remove_info := ServiceRemoveInfo{}
		json.Unmarshal([]byte(detail), &service_remove_info)
		log.Println("[Result]:")
		log.Println(service_remove_info)

		if err := collection.Insert(service_remove_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	}else if item.ActionName == "search_performance_info"{
		search_performance_info := SearchPerformanceInfo{}
		json.Unmarshal([]byte(detail), &search_performance_info)
		log.Println("[Result]:")
		log.Println(search_performance_info)

		if err := collection.Insert(search_performance_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	}else if item.ActionName == "ni_performance_info"{
		ni_performance_info := NiPerformanceInfo{}
		json.Unmarshal([]byte(detail), &ni_performance_info)
		log.Println("[Result]:")
		log.Println(ni_performance_info)

		if err := collection.Insert(ni_performance_info); err != nil {
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