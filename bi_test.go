package misbi

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/MiSingularity/logging/p"
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
	resp, err := http.Post("http://127.0.0.1:8088/bi?action=deepshare/userlink", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
