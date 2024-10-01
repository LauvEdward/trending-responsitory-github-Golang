package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"trending-github-golang/db"
	"trending-github-golang/handler"
	"trending-github-golang/log"
	"trending-github-golang/responsitory/repo_impl"
	"trending-github-golang/router"
)

func init() {
	os.Setenv("APP_NAME", "trending-github-golang")
	os.Setenv("APP_VERSION", "v0.0.1")
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
	userRepo := repo_impl.NewUserRepoImpl(sql)
	userHandle := handler.UserHandler{
		UserRepoImpl: userRepo,
	}
	api := router.API{
		Echo:        e,
		UserHandler: &userHandle,
	}
	api.SetUp()
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	println("Hello, World!")
	return c.String(http.StatusOK, "Hello, World!")
}
