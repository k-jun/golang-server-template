package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

type Config struct {
	Psql psqlConfig
}

type psqlConfig struct {
	Dbname string `toml:"dbname"`
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

var DB *sql.DB
var CTX context.Context

func init() {
	boil.DebugMode = true
	var err error
	var config Config
	_, err = toml.DecodeFile("./sqlboiler.toml", &config)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Psql.Host, config.Psql.Port, config.Psql.User, config.Psql.Pass, config.Psql.Dbname,
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Could not connect to database")
		// defer DB.Close()
		panic(err)
	}
	CTX = context.Background()
}
