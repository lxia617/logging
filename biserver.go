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

var (
	projBiLogs map[string][]*p.BiLog = make(map[string][]*p.BiLog)
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
	pojName := pbBiLog.ProjectName

	if biLogs, ok := projBiLogs[pojName]; !ok {
		projBiLogs[pojName] = []*p.BiLog{pbBiLog}
	} else {
		biLogs = append(biLogs, pbBiLog)
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
			ret = append(ret, []byte(bilog.String()))
		}
	}
	return
}
