package misbi

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"html"
	"io/ioutil"
	"log"
	"misbi/p"
	"net/http"
	"strconv"
)

type BILog struct {
	ProjectName string
	ActionName  string
	Timestamp   int64
	Data        []byte
}

var (
	projBiLogs map[string][]*BILog = make(map[string][]*BILog)
)

func BiFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello BI: %s", html.EscapeString(req.URL.Path))

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Bi Body:", body)
	pbBiLog := &p.BiLog{}
	if err := proto.Unmarshal(body, pbBiLog); err != nil {
		log.Panic(err)
	}
	log.Println(pbBiLog.String())

	log.Println(pbBiLog.ExtensionMap())

	//Save the log by projectName
	appendBiLog(pbBiLog)

}

func GetBiLogsFunc(w http.ResponseWriter, req *http.Request) {
	projName := req.FormValue("project")
	startTime, _ := strconv.Atoi(req.FormValue("from"))
	endTime, _ := strconv.Atoi(req.FormValue("to"))
	log.Println(startTime, endTime)
	w.Write(GetBiLogsForProj(projName)[0])

}

func appendBiLog(pbBiLog *p.BiLog) {
	b, err := proto.Marshal(pbBiLog)
	if err != nil {
		panic(err)
	}
	pojName := pbBiLog.GetProjectName()
	biLog := &BILog{
		ProjectName: pojName,
		ActionName:  pbBiLog.GetActionName(),
		Timestamp:   pbBiLog.GetTimestamp(),
		Data:        b,
	}
	if biLogs, ok := projBiLogs[pojName]; !ok {
		projBiLogs[pojName] = []*BILog{biLog}
	} else {
		biLogs = append(biLogs, biLog)
	}
	log.Println("Project bi logs:", projBiLogs)
}

func GetBiLogsForProj(projName string) (ret [][]byte) {
	if projName == "" {
		log.Printf("Invalid request, should have a form: name=<projectName>\n")
	} else if bilogs, ok := projBiLogs[projName]; !ok {
		log.Printf("No Logs for project:%s\n", projName)
	} else {
		for i, bilog := range bilogs {
			log.Printf("	[%d] %#v\n", i, bilog)
			ret = append(ret, bilog.Data)
		}
	}
	return
}
