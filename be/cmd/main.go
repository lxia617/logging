package main

import (
	"github.com/MiSingularity/logging/be"
)

const (
	BI_SERVER_PORT = "8999"
)

func main() {
	be.InitDbConn()
	be.InitGrpcServer(BI_SERVER_PORT)
}
