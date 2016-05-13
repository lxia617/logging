package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"strings"
	"github.com/MISingularity/logging/be"
	"log"
	"encoding/json"
	"time"
)
const (
	mongoHost = "42.159.133.35"
	mongoPort = "27017"
	MONGO_DBNAME = "XIAOZHI_LOG"
)

var DB *mgo.Database

func main() {
	if err := be.InitDbConn(mongoHost,mongoPort); err != nil {
		log.Fatal(err)
	}

	DB = be.MgoSession.DB(MONGO_DBNAME)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "collections:\n" + strings.Join(showAllCollections(),"\n"))
	})
	router.GET("/collection/:name", func(c *gin.Context) {
		name, _ := c.Params.Get("name")
		c.String(http.StatusOK, showCollectionData(name),"\n")
	})
	router.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusUnauthorized, "not authorized")
	})
	router.PUT("/error", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "an error happened :(")
	})
	router.Run(":8080")
}

func showAllCollections() (names []string) {
	names, _ = DB.CollectionNames()

	//var links = make(string, len(names))

/*	for i, name := range names {
		names[i] = "<a href='localhost:8080/collection?name=" + name +"'>" + name +"</a>";
	}

	if err != nil {
		log.Println("[ERROR]Can not connect collections " + MONGO_DBNAME +", err:", err)
	}*/

	return names
}

func showCollectionData(name string) (result string) {
	collection := DB.C(name)

	var err error
	if name ==  "service_start_info"{
		var serviceStartInfos []be.ServiceStartInfo
		err = collection.Find(nil).All(&serviceStartInfos)
		for _, serviceStartInfo := range serviceStartInfos {
			b, _ := json.MarshalIndent(serviceStartInfo, "\t", "")
			result += string(b) + "\n"
		}
	}else if name == "action_perform_result"{
		var actionPerformResults []be.ActionPerformResult
		err = collection.Find(nil).All(&actionPerformResults)

		for _, actionPerformResult := range actionPerformResults {
			queryTime := time.Unix(actionPerformResult.QueryTimestamp/1000, actionPerformResult.QueryTimestamp%1000)
			result += queryTime.String() + "\n"
			b, _ := json.MarshalIndent(actionPerformResult, "", "")
			result += string(b) + "\n\n"
		}
	}

	//err := collection.Find(nil).All(&results)
	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
	}
	return result
}

