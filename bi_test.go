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
		ProjectName: proto.String("deepshare"),
		ActionName:  proto.String("userlink"),
		Timestamp:   proto.Int64(time.Now().Unix()),
	}

	if err := proto.SetExtension(&pbBiLog, p.E_DeepshareLog, &p.DeepshareBiLog{
		Os:        proto.String("ios"),
		OsVersion: proto.String("8.3"),
		DeviceId:  proto.String("11111"),
		Ip:        proto.String("127.0.0.1"),
		Param:     proto.String("0000009999"),
	}); err != nil {
		log.Println("!!!!!!!!error when set extension", err)
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
		ext, err := proto.GetExtension(pbBiLog, p.E_DeepshareLog)
		if err != nil {
			panic(err)
		}
		switch ext.(type) {
		case *p.DeepshareBiLog:
			deepshareLog := ext.(*p.DeepshareBiLog)
			log.Printf("deepshare log: %#v\n", deepshareLog)
			log.Println("deepshare log in string format: ", deepshareLog.String())
		}

	}

}
