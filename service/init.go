package service

import (
	"github.com/go-xorm/xorm"
	"github.com/guoyueqiang2013/goim/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBEngin *xorm.Engine
	DBErr   error
)

func init() {
	dn := "mysql"
	dsn := "root:111111@(192.168.1.188:3306)/chat?charset=utf8"
	DBEngin, DBErr = xorm.NewEngine(dn, dsn)
	if DBErr != nil {
		log.Fatal(DBErr.Error())
	}
	if err := DBEngin.Ping(); err != nil {
		log.Fatal("Database Engine Ping() error")
	}
	log.Printf("The database initing succ!!\n")

	DBEngin.ShowSQL(true)
	DBEngin.SetMaxOpenConns(10)

	//自动建表
	DBEngin.Sync2(new(model.User),new(model.Contact),new(model.Community))
}