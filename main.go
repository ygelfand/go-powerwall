package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ygelfand/go-powerwall/internal/api"
	"github.com/ygelfand/go-powerwall/internal/powerwall"
)

func main() {
	pwr := powerwall.NewPowerwallGateway()
	app := api.NewApi(pwr)
	go pwr.PeriodicRefresh(15 * time.Second)

	router := gin.Default()
	base := router.Group("/api")
	{
		v1 := base.Group("/v1")
		{
			v1.GET("/strings", app.Strings)
		}
	}
	x := pwr.RunQuery("DeviceControllerQuery", nil)
	fmt.Println("start")
	var prettyJSON bytes.Buffer
	/*x := pwr.RunQuery("ComponentsQuery", nil)
	err := json.Indent(&prettyJSON, []byte(x), "", "\t")
	if err != nil {
		fmt.Println("JSON parse error: ", err)
	}
	//log.Println(string(prettyJSON.Bytes()))*/
	x = pwr.GetConfig()
	err := json.Indent(&prettyJSON, []byte(x), "", "\t")
	if err != nil {
		fmt.Println("JSON parse error: ", err)
	}
	log.Println(string(prettyJSON.Bytes()))
	router.Run("localhost:8080")
}
