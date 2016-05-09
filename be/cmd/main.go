package main

import (
	"flag"
	"log"
	"os"

	"github.com/MISingularity/logging/be"
	"encoding/json"
	"fmt"
)

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func testResponse() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}

type Action struct {
	Command string `json:"command"`
	Actions string `json:"actions"`
}

func testAct() {
	str := `{"command":"commandName", "actions": "acts1"}`
	res := Action{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Command)
}

func main() {
	testAct()
	fs := flag.NewFlagSet("dslogging", flag.ExitOnError)
	port := fs.String("port", "50051", "Listening port")
	mh := fs.String("mongo-host", "127.0.0.1", "MongoDB host")
	mp := fs.String("mongo-port", "27017", "MongoDB port")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	if fi, err := os.Stat("log"); os.IsNotExist(err) || !fi.IsDir() {
		os.Mkdir("log", 0755)
	}

	if err := be.InitDbConn(*mh, *mp); err != nil {
		log.Fatal(err)
	}
	be.InitGrpcServer(*port)
}
