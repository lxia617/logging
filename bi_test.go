package misbi

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"misbi/p"
	"net/http"
	"testing"
	"time"
)

func TestBiDeepshareLog(t *testing.T) {

	pbBiLog := p.BiLog{
		ProjectName: "deepshare",
		ActionName:  "userlink",
		Timestamp:   time.Now().Unix(),
		Detail:      []byte("deepshare bi detail"),
	}

	log.Println("Sending bi request with deepsharelog to logging server", pbBiLog)

	b, err := proto.Marshal(&pbBiLog)
	if err != nil {
		log.Panic(err)
	}
	resp, err := http.Post("http://127.0.0.1:8088/bi", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func TestGetBiLogsByProjName(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8088/logs?project=deepshare")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(body)

	pbBiLog := &p.BiLog{}
	if err := proto.Unmarshal(body, pbBiLog); err != nil {
		panic(err)
	} else {
		log.Println("~~~~~~~~", pbBiLog.String())
	}

}
