package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var CTX context.Context

func Initialize() error {
	var err error

	// TODO 環境変数
	dsn := "dena:dena@tcp(127.0.0.1:3306)/sample"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	CTX = context.Background()
	return nil
}
