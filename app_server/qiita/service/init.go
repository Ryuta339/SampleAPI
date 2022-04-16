package service

import (
	"errors"
	"log"
	"qiita/model"

	"github.com/go-xorm/xorm"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

var DbEngine *xorm.Engine

func init() {
	// driverName := "sqlite3"
	// DsName := "./database/test.db"
	driverName := "mysql"
	DsName := "root:password123@tcp(arch_db:3306)/books?charset=utf8"
	err := errors.New("")
	// := を使うとDbEngineがローカルに束縛される
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)

	sampleBook := model.Book{Title: "Programing Go", Content: "foobar"}
	err = DbEngine.Sync2(&sampleBook)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	log.Println("init data base ok")
}
