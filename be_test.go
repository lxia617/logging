package logging

import (
	"github.com/MISingularity/logging/be"
	"github.com/MISingularity/logging/be/internal"
	"github.com/MISingularity/logging/p"
	"testing"
)

func TestSaveBiLog(t *testing.T) {
	be.InitDbConn("127.0.0.1", "27017")
	pbBiLog := &p.BiLog{
		ProjectName: "testprojectXXX",
		ActionName:  "testactionXXX",
		Timestamp:   int64(88888),
		Detail:      []byte("testing bi detail"),
	}
	be.SaveBiLog(pbBiLog)

	retriedvedLog := internal.RetrieveBiLog(pbBiLog)
	if !internal.IsEqual(retriedvedLog, pbBiLog) {
		t.Error("save and retrieve not agree")
	}

	internal.DeleteBiLog(pbBiLog)
}
