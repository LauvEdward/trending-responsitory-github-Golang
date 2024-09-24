package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"runtime"
	"trending-github-golang/db"
	"trending-github-golang/handler"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "1234",
		DbName:   "postgres",
	}
	logError("Error connecting to database")
	sql.Connect()
	defer sql.Close() // after main work
	e := echo.New()
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":3000"))
}

func logError(err string) {
	_, file, line, _ := runtime.Caller(1)
	logrus.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Fatal(err)
}
