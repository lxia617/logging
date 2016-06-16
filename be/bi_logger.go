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
	today := time.Now().Format("20060102")
	if l := loggers[projName + today]; l != nil {
		return l
	}
	logfile, err := os.OpenFile("log/" + projName + "_" + time.Now().Format("20060102.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open file: ", err)
		return log.New(os.Stdout, "["+projName+"]", log.LstdFlags)
	}
	l := log.New(io.MultiWriter(logfile, os.Stdout), "["+projName+"]", log.LstdFlags)
	loggers[projName + today] = l
	return l
}

// the difference between this log and writeLog is this log is:
// writing all the projectName to file, but writeLog write project related log to corresponding file
// the format of this log file may needs to be changed to be parsed by logstash
// this log file seems to be duplicated with some log in the <projName_time.log> file
func SetLogFile() {
	f, err :=  os.OpenFile("log/" + "NIAllLog_" + time.Now().Format("20060102.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
}