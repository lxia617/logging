package be

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func ShowDataInBrowser() {
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
	sessionCopy := MgoSession.Copy()
	defer sessionCopy.Close()
	log.Print("get / request, showAllCollections")
	log.Print(time.Now())

	names, err := sessionCopy.DB(MONGO_DBNAME).CollectionNames()

	if err != nil {
		log.Println("[ERROR]Can not connect Database "+MONGO_DBNAME+", err:", err)

		c.String(http.StatusGatewayTimeout, err.Error())
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
	sessionCopy := MgoSession.Copy()
	defer sessionCopy.Close()
	log.Print("get showCollectionData request:" + name + ", pageId:" + strconv.Itoa(pageId) + ", success: " + success)
	log.Print(time.Now())

	collection := sessionCopy.DB(MONGO_DBNAME).C(name)

	if collection == nil {
		log.Printf("Collection nil, maybe error")
		c.String(http.StatusGatewayTimeout, "Collection nil, please try <host:port>/ to see whether this Collection is in the DB")
		return
	}

	var err error
	var result string

	if name == "service_start_info" {
		var serviceStartInfos []ServiceStartInfo
		err = collection.Find(nil).All(&serviceStartInfos)

		for _, serviceStartInfo := range serviceStartInfos {
			b, _ := json.MarshalIndent(serviceStartInfo, "\t", "")
			result += string(b) + "\n"
		}
	} else if name == "action_perform_result" {

		query := bson.M{"success": strings.Compare(success, "true") == 0}
		log.Println(query)

		var actionPerformResults []ActionPerformResult
		//err = collection.Find(bson.M{"success":false}).Sort("-querytimestamp").All(&actionPerformResults)
		allData := collection.Find(query).Sort("-querytimestamp")
		count, _ := allData.Count()

		err = allData.Skip(pageSize * (pageId - 1)).Limit(pageSize).All(&actionPerformResults)

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

		c.Writer.Write([]byte(linkpage + "<br/><br/>"))
		result = strings.Replace(result, "\n", "<br/>", -1)
	} else {
		var BiLogStrInfos []BiLogStr

		if err = collection.Find(nil).Sort("-timestamp").All(&BiLogStrInfos); err == nil {
			for _, logStr := range BiLogStrInfos {
				result += logStr.Detail + "\n"
				break
			}
		}
	}

	//err := collection.Find(nil).All(&results)
	if err != nil {
		log.Printf("RunQuery showCollectionData : ERROR : %s\n, please try again", err)
		c.String(http.StatusGatewayTimeout, err.Error())
	}

	log.Print(result)
	c.String(http.StatusOK, result)
}
