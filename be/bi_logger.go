package be

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/MISingularity/logging/p"
)

var loggers map[string]*log.Logger = make(map[string]*log.Logger)

func writeLog(biLog *p.BiLog) {
	logger := GetLogger(biLog.ProjectName)
	logger.Println(biLog.ProjectName, biLog.ActionName, biLog.Timestamp, string(biLog.Detail))
}

func GetLogger(projName string) *log.Logger {
	if l := loggers[projName]; l != nil {
		return l
	}
	logfile, err := os.OpenFile("log/" + projName + "_" + time.Now().Format("20060102_150405.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open file: ", err)
		return log.New(os.Stdout, "["+projName+"]", log.LstdFlags)
	}
	l := log.New(io.MultiWriter(logfile, os.Stdout), "["+projName+"]", log.LstdFlags)
	loggers[projName] = l
	return l
}

func SetLogFile() {
	f, err :=  os.OpenFile("log/" + "be_main_" + time.Now().Format("20060102_15.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
}