package main

import (
	"log"
	"misbi"
	"net/http"
)

func main() {
	misbi.InitDbConn()
	go func() {
		log.Println("start http server")
		http.HandleFunc("/bi", misbi.BiFunc)
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Fatal(err)
		}
	}()

	misbi.InitGrpcServer()
}
