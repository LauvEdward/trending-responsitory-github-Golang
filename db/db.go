package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.Username, s.Password, s.DbName)
	println("data source" + dataSource)
	s.Db = sqlx.MustConnect("postgres", dataSource)
	err := s.Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	println("DB connected")
}

func (s *Sql) Close() {
	if s.Db != nil {
		err := s.Db.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	println("DB disconnected")
}
