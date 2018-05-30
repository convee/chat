package mysql

import (
	"chat/conf"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(name string) (db *sqlx.DB) {
	config := conf.Get().Mysql[name]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Ip, config.Port, config.DbName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return
}
