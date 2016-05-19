package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MISingularity/logging/be"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoHost    = "42.159.133.35"
	mongoPort    = "27017"
	MONGO_DBNAME = "XIAOZHI_LOG"
)

func main() {
	if err := be.InitDbConn(mongoHost, mongoPort); err != nil {
		log.Fatal(err)
	}

	be.SetLogFile()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, strings.Join(showAllCollections(),"\n"))
		showAllCollections(c)
	})
	router.GET("/collection/:name", func(c *gin.Context) {
		name, _ := c.Params.Get("name")
		pageId, _ := strconv.Atoi(c.DefaultQuery("pageId", "1"))
		pageid, _ := strconv.Atoi(c.DefaultQuery("pageid", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

		if pageId == 1 {
			pageId = pageid
		}

		success := c.DefaultQuery("success", "false")
		showCollectionData(c, name, pageId, success, pageSize)
		//c.String(http.StatusOK, showCollectionData(name),"\n")
	})
	router.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusUnauthorized, "not authorized")
	})
	router.PUT("/error", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "an error happened :(")
	})

	router.Run(":8080")
}

func showAllCollections(c *gin.Context) {
	log.Print("get / request, showAllCollections")
	log.Print(time.Now())

	names, err := be.MgoSession.DB(MONGO_DBNAME).CollectionNames()

	if err != nil {
		log.Println("[ERROR]Can not connect Database "+MONGO_DBNAME+", err:", err)

		c.String(http.StatusGatewayTimeout, err.Error())

		if errConnect := be.InitDbConn(mongoHost, mongoPort); errConnect != nil {
			log.Fatal(errConnect)
			c.String(http.StatusGatewayTimeout, "try to connect again fail")
			c.String(http.StatusGatewayTimeout, err.Error())
			return
		} else {
			c.String(http.StatusRequestTimeout, "Please try again")
		}
	}

	// if use something like this, will not show link, just show raw data
	log.Print(names)

	links := make([]string, len(names))

	for i, name := range names {
		//links[i] = "<a href='http://localhost:8080/collection/" + name +"'>" + name +"</a><br/>";
		links[i] = "<a href='/collection/" + name + "'>" + name + "</a><br/>"
	}

	//c.String(http.StatusOK, strings.Join(links,"\n"))
	c.Writer.Write([]byte(strings.Join(links, "\n")))

}

func showCollectionData(c *gin.Context, name string, pageId int, success string, pageSize int) {
	log.Print("get showCollectionData request:" + name + ", pageId:" + strconv.Itoa(pageId) + ", success: " + success)
	log.Print(time.Now())

	collection := be.MgoSession.DB(MONGO_DBNAME).C(name)

	if collection == nil {
		log.Printf("Collection nil, maybe error")
		c.String(http.StatusGatewayTimeout, "Collection nil, please try <host:port>/ to see whether this Collection is in the DB")
		return
	}

	var err error
	var result string

	if name == "service_start_info" {
		var serviceStartInfos []be.ServiceStartInfo
		err = collection.Find(nil).All(&serviceStartInfos)

		for _, serviceStartInfo := range serviceStartInfos {
			b, _ := json.MarshalIndent(serviceStartInfo, "\t", "")
			result += string(b) + "\n"
		}
	} else if name == "action_perform_result" {

		query := bson.M{"success": strings.Compare(success, "true") == 0}
		log.Println(query)

		var actionPerformResults []be.ActionPerformResult
		//err = collection.Find(bson.M{"success":false}).Sort("-querytimestamp").All(&actionPerformResults)
		allData := collection.Find(query).Sort("-querytimestamp")
		count, _ := allData.Count()

		err = allData.Skip(pageSize * (pageId - 1)).Limit(pageSize).All(&actionPerformResults)

		//err := collection.Find(nil).All(&results)
		for i := 1; err != nil && i < 5; i++ {
			log.Printf("Action Perform Result : ERROR : %s\n", err)
			log.Println("Try to connect times:%d", i)

			if errConnect := be.InitDbConn(mongoHost, mongoPort); errConnect != nil {
				log.Fatal(errConnect)
			}

			err = collection.Find(query).Sort("-querytimestamp").Skip(pageSize * (pageId - 1)).Limit(pageSize).All(&actionPerformResults)
		}

		for index, actionPerformResult := range actionPerformResults {
			queryTime := time.Unix(actionPerformResult.QueryTimestamp/1000, actionPerformResult.QueryTimestamp%1000)
			result += "Record Index:" + strconv.Itoa(index+1+pageSize*(pageId-1)) + "\n"
			result += queryTime.String() + "\n"
			//result += strconv.FormatInt(actionPerformResult.QueryTimestamp, 10) + "\n"
			b, _ := json.MarshalIndent(actionPerformResult, "", "")
			result += string(b) + "\n\n---------------------------------------------------------------------------------------------------------------\n\n"
		}

		var linkpage string

		pageIndex := 1
		for pageIndex <= count/pageSize+1 {
			linkpage += "<a href='/collection/action_perform_result?pageid=" + strconv.Itoa(pageIndex) + "&pagesize=" + strconv.Itoa(pageSize) + "&success=" + success + "'>" + strconv.Itoa(pageIndex) + "</a>   "
			pageIndex = pageIndex + 1
		}

		result = strings.Replace(result, "\n", "<br/>", -1)
	} else {
		var BiLogStrInfos []be.BiLogStr
		err = collection.Find(nil).All(&BiLogStrInfos)

		for _, logStr := range BiLogStrInfos {
			b, _ := json.MarshalIndent(logStr, "", "")
			result += string(b) + "<br/><br/>"
		}
	}

	//err := collection.Find(nil).All(&results)
	if err != nil {
		log.Printf("RunQuery showCollectionData : ERROR : %s\n, please try again", err)

		if errConnect := be.InitDbConn(mongoHost, mongoPort); errConnect != nil {
			log.Fatal(errConnect)
		}
		c.String(http.StatusGatewayTimeout, err.Error())
	}

	log.Print(result)
	c.String(http.StatusOK, result)
}
