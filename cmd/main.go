package main

import (
	"log"
	"misbi"
	"net/http"
)

func main() {
	http.HandleFunc("/bi", misbi.BiFunc)
	http.HandleFunc("/logs", misbi.GetBiLogsFunc)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
