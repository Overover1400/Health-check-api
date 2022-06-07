package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/overover1400/api-health-check/adaptor/store"
	"github.com/overover1400/api-health-check/delivery/http/v1"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

func main() {

	//TODO use "swagger" for document
	db := "root:root@tcp(localhost:3306)/api?charset=utf8mb4&parseTime=True&loc=Local"
	//db := "root:root@tcp(health_check_api_db:3306)/api?charset=utf8mb4&parseTime=True&loc=Local"
	//	os.Getenv("MYSQL_USER"),
	//	os.Getenv("MYSQL_PASSWORD"),
	//	os.Getenv("MYSQL_HOST"),
	//	os.Getenv("MYSQL_PORT"),
	//	os.Getenv("MYSQL_DATABASE"))

	//connect to databases and auto migrate
	mysql := store.New(db)

	go func() {
		c := cron.New()

		// every 1 minute select from databases to send health check api request
		err := c.AddFunc("1 * * * *", func() {
			// cron will check databases every minute

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

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello from test")
	})

	//TODO Gracefully shutdown with os.Sig
	e.Logger.Fatal(e.Start(":8080"))

}
