package be

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MISingularity/logging/p"
	"gopkg.in/mgo.v2"
)

const (
	MONGO_USER     = ""
	MONGO_PASSWORD = ""
	MONGO_DBNAME = "XIAOZHI_LOG"
)

type BiLogStr struct {
	ProjectName string `json:"projectName,omitempty"`
	ActionName  string `json:"actionName,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	Detail      string `json:"detail,omitempty"`
}

type ServiceStartInfo struct {
	DeviceId    string `json:"deviceId"`
	NiVersion   string `json:"niVersion"`
	Model       string `json:"model"`
	Manufacture string `json:"manufacture"`
	Fingerprint string `json:"fingerprint"`
}

type ServiceRemoveInfo struct {
	DeviceId string `json:"deviceId"`
}

type NoSearchResultInfo struct {
	DeviceId       string `json:"deviceId"`
	QueryTimestamp int64  `json:"queryTimestamp"`
	Command        string `json:"command"`
}

type ActionPerformResult struct {
	DeviceId                    string `json:"deviceId"`
	QueryTimestamp              int64  `json:"queryTimestamp"`
	Success                     bool   `json:"success"`
	CurrentQuery                string `json:"currentQuery"`
	QueryTitle                  string `json:"queryTitle"`
	OpenPackageName             string `json:"openPackageName"`
	OpenPackageVersion          string `json:"openPackageVersion"`
	CanPerformPathCount         int32  `json:"canPerformPathCount"`
	ActionListCount             int32  `json:"actionListCount"`
	ActionList                  string `json:"actionList"`
	ReceiveQueryResultTimestamp int64  `json:"receiveQueryResultTimestamp"`
}

type NiPerformanceInfo struct {
	DeviceId                    string `json:"deviceId"`
	QueryTimestamp              int64  `json:"queryTimestamp"`
	QueryVoiceTimestamp         int64  `json:"queryVoiceTimestamp"`
	ReceiveVoiceResultTimestamp int64  `json:"receiveVoiceResultTimestamp"`
	SendQueryTimestamp          int64  `json:"sendQueryTimestamp"`
	ReceiveQueryResultTimestamp int64  `json:"receiveQueryResultTimestamp"`
}

type Action struct {
}

var MgoSession *mgo.Session

var mongoHost string
var mongoPort string

func InitDbConn(_mongoHost, _mongoPort string) error {
	mongoHost = _mongoHost
	mongoPort = _mongoPort
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
	sessionCopy := MgoSession.Copy()
	defer sessionCopy.Close()
	if MgoSession == nil {
		log.Println("MongoDB not connected, user log saved to file")
		log.Printf("[UserLog] %#v\n", item)
		panic(MgoSession)
		return errors.New("Can not connect mongoDb")
	}
	db := sessionCopy.DB(item.ProjectName)
	collection := db.C(item.ActionName)

	detail := string(item.Detail[:])
	detail = strings.Replace(detail, "'", "\"", -1)

	itemStr := &BiLogStr{
		ProjectName: item.ProjectName,
		ActionName:  item.ActionName,
		Timestamp:   item.Timestamp,
		Detail:      detail,
	}

	log.Printf("[UserLog] %#v\n", itemStr)

	if item.ActionName == "service_start_info" {
		service_start_info := ServiceStartInfo{}
		json.Unmarshal([]byte(detail), &service_start_info)
		log.Println("[Result]:")
		log.Println(service_start_info)

		if err := collection.Insert(service_start_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	} else if item.ActionName == "service_remove_info" {
		service_remove_info := ServiceRemoveInfo{}
		json.Unmarshal([]byte(detail), &service_remove_info)
		log.Println("[Result]:")
		log.Println(service_remove_info)

		if err := collection.Insert(service_remove_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	} else if item.ActionName == "no_search_result" {
		no_search_result := NoSearchResultInfo{}
		json.Unmarshal([]byte(detail), &no_search_result)
		log.Println("[Result]:")
		log.Println(no_search_result)

		if err := collection.Insert(no_search_result); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	} else if item.ActionName == "action_perform_result" {
		search_performance_info := ActionPerformResult{}
		json.Unmarshal([]byte(detail), &search_performance_info)
		log.Println("[Result]:")
		log.Println(search_performance_info)

		if err := collection.Insert(search_performance_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	} else if item.ActionName == "ni_performance_info" {
		ni_performance_info := NiPerformanceInfo{}
		json.Unmarshal([]byte(detail), &ni_performance_info)
		log.Println("[Result]:")
		log.Println(ni_performance_info)

		if err := collection.Insert(ni_performance_info); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	} else if item.ActionName == "tracking_msg" {
		if err := collection.Insert(item); err != nil {
			log.Println("[ERROR]Save user log to MongoDB failed, err:", err)
			return err
		}
	}
	return nil
}
