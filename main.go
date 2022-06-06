package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/overover1400/api-health-check/adaptor/store"
	"github.com/overover1400/api-health-check/delivery/http/v1"
	"github.com/robfig/cron"
	"log"
	"os"
)

func main() {

	//TODO use "swagger" for document
	//dataBaseName := "client_api_user:client_api_pass@tcp(127.0.0.1:3306)/client_api?charset=utf8mb4&parseTime=True&loc=Local"
	db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	//connect to databases and auto migrate
	mysql := store.New(db)

	go func() {
		c := cron.New()

		// every 1 minute select from databases to send health check api request
		err := c.AddFunc("1 * * * *", func() {
			// cron will check databases every minute
			//TODO to send request internally should communicate with Protocol Buffers

			if err := v1.HealthCheckApi(mysql); err != nil {
				log.Fatalln("health Check api has error : ", err)
			}

			fmt.Println("Run Cron every 1 minute")

		})
		if err != nil {
			log.Fatalln("cron error : ", err)
			//c.Stop()
		}

		c.Start()
	}()

	// setup http server and router
	e := echo.New()

	// register API with check time interval
	e.POST("/register", v1.CreateClientApi(mysql))

	e.PUT("/toggle/:id", v1.ToggleClientApi(mysql))

	// get list of our API and it is CDN-Friendly
	e.GET("/list", v1.FindClientApis(mysql))

	e.DELETE("/delete/:id", v1.DeleteClientApi(mysql))

	//TODO Gracefully shutdown with os.Sig
	e.Logger.Fatal(e.Start(":8080"))

}
