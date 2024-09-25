package main

import (
	"github.com/labstack/echo/v4"
	"os"
	"trending-github-golang/db"
	"trending-github-golang/handler"
	"trending-github-golang/log"
	"trending-github-golang/router"
)

func init() {
	os.Setenv("APP_NAME", "trending-github-golang")
	//os.Setenv("APP_VERSION", "v0.0.1")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "1234",
		DbName:   "postgres",
	}
	sql.Connect()
	defer sql.Close() // after main work

	e := echo.New()
	userHandle := handler.UserHandler{}
	api := router.API{
		Echo:        e,
		UserHandler: &userHandle,
	}
	api.SetUp()
	e.Logger.Fatal(e.Start(":3000"))
}
