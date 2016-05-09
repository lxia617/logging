package main

import (
	"log"
	"time"

	"strconv"

	"flag"

	"os"

	"fmt"

	"github.com/MISingularity/logging/fe"
	"github.com/MISingularity/logging/p"
)

func main() {
	log.Println("init bi client")

	fs := flag.NewFlagSet("logging_test", flag.ContinueOnError)
	serverAddr := fs.String("server-addr", "127.0.0.1:50051", "Specify the server url to connect to")
	runs := fs.Int("runs", 2, "How many loops")
	interval := fs.Int("interval", 0, "Interval between two grpc call(in Seconds)")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Panic(err)
	}
	log.Println("arguments are:")
	fs.VisitAll(func(flag *flag.Flag) {
		fmt.Println("			", flag.Name, ":", flag.Value)
	})
	if err := fe.Init(*serverAddr); err != nil {
		log.Panic(err)
	}
	for i := 0; i < *runs; i++ {
		log.Println("--------------------", i)
/*		log.Println("add a bi log for deepshare")
		fe.Bi(&p.BiLog{
			ProjectName: "deepshare",
			ActionName:  "userlink",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("detail~~~~~~" + strconv.Itoa(i)),
		})
		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("add a bi log for ime")
		fe.Bi(&p.BiLog{
			ProjectName: "ime",
			ActionName:  "input",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("detail~~~~~~" + strconv.Itoa(i)),
		})*/

		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("ni service add & removed log")
		fe.Bi(&p.BiLog{
			ProjectName: "ni",
			ActionName:  "service_log",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("service removed" + strconv.Itoa(i)),
		})

		time.Sleep(time.Second * time.Duration(*interval))
		fe.Bi(&p.BiLog{
			ProjectName: "ni",
			ActionName:  "service_log",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("service restarted" + strconv.Itoa(i)),
		})

		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("ni performance log")
		fe.Bi(&p.BiLog{
			ProjectName: "ni",
			ActionName:  "ni_performance",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("action no result" + strconv.Itoa(i)),
		})


		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("ni performance log")
		fe.Bi(&p.BiLog{
			ProjectName: "ni",
			ActionName:  "ni_performance",
			Timestamp:   time.Now().Unix(),
			Detail:      []byte("current query path perform success query title packageName" + strconv.Itoa(i)), // need to analyse the action list format and see whether we can parse them
		})

		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("ni tracking data")

		fe.Bi(&p.BiLog{
			ProjectName: "ni",
			ActionName:  "tracking_data",
			Timestamp:   time.Now().Unix(),
			//Detail:      []byte("{\"command\":\"" + "commandName" + "\", \"actions\": \"" + "acts" + strconv.Itoa(i) + "\"}"),
			//Detail:      []byte("{\"command\":\"" + "commandName" + "\", \"actions\": \"" + "acts" + strconv.Itoa(i) + "\"}"),
			Detail:      []byte("{'command':'" + "commandName" + "', 'actions': '" + "acts" + strconv.Itoa(i) + "'}"), // need to analyse the action list format and see whether we can parse them
	})

		time.Sleep(time.Second * time.Duration(*interval))
		log.Println("add a DeviceInfo log for ni")
		fe.DeviceInfo(&p.DeviceInfo{
			DeviceId: "DeviceId",
			Manufacturer:  "Manufacturer",
			OpenAppTime:   time.Now().Unix(),
			Model: "{'command':'" + "commandName" + "', 'actions': '" + "acts" + strconv.Itoa(i) + "'}",
			NiVersion: "niversion",
			AppVersion: "AppVersion",
		})
		time.Sleep(time.Second * time.Duration(*interval))
	}

}
