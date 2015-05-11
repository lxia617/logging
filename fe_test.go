package logging

import (
	"github.com/MiSingularity/logging/be/internal"
	"github.com/MiSingularity/logging/fe"
	"github.com/MiSingularity/logging/p"
	"log"
	"os"
	"testing"
)

func setup() {
	fe.Init("127.0.0.1", "8999")
}

func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func TestGrpcCliCall(t *testing.T) {

	pbBiLog := &p.BiLog{
		ProjectName: "testprojectXXX",
		ActionName:  "testactionXXX",
		Timestamp:   int64(88888),
		Detail:      []byte("testing bi detail"),
	}

	log.Println("Sending bi request with deepsharelog to logging server", pbBiLog)
	if err := fe.Bi(pbBiLog); err != nil {
		t.Error("fe.Bi() failed, err:", err)
	}

	internal.DeleteBiLog(pbBiLog)

}
