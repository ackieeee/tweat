package adapter

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Tweat() *sqlx.DB {
	host := os.Getenv("MYSQL_TWEAT_HOST")
	user := os.Getenv("MYSQL_TWEAT_USER")
	password := os.Getenv("MYSQL_TWEAT_PASSWORD")
	port := os.Getenv("MYSQL_TWEAT_PORT")
	dbName := os.Getenv("MYSQL_TWEAT_DATABASE")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, host, port, dbName)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
