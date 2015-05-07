package misbi

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"html"
	"io/ioutil"
	"log"
	"misbi/p"
	"net/http"
)

var (
	projBiLogs map[string][]*p.BiLog = make(map[string][]*p.BiLog)
)

func BiFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello BI: %s", html.EscapeString(req.URL.Path))

	action := req.FormValue("action")
	log.Println("---------------Action: ", action)

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

	//Save the log
	SaveBiLog(pbBiLog)

}
