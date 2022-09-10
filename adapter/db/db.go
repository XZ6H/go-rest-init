package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"myapp/config"
)

func New(conf *config.DbConf) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.DbName,
		conf.Port,
	)
	return sql.Open("postgres", dsn)
}
