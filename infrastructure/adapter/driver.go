package adapter

import (
	"errors"
	"log"

	_ "github.com/lib/pq"
	"rizhua.com/infrastructure/etc"
	"xorm.io/xorm"
)

func NewDb(driver string) (engine *xorm.Engine, err error) {
	switch driver {
	case "postgres":
		engine, err = newPostgres()
	default:
		err = errors.New("未知的数据库类型")
	}
	return
}

func newPostgres() (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine("postgres", etc.C.Postgres.DSN())
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	engine.ShowSQL(true)
	return
}
