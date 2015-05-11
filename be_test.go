package logging

import (
	"github.com/MiSingularity/logging/be"
	"github.com/MiSingularity/logging/be/internal"
	"github.com/MiSingularity/logging/p"
	"testing"
)

func TestSaveBiLog(t *testing.T) {
	be.InitDbConn()
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
