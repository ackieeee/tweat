package adapter

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	gm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Tweat() *sqlx.DB {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("MYSQL_TWEAT_USER"),
		os.Getenv("MYSQL_TWEAT_PASSWORD"),
		os.Getenv("MYSQL_TWEAT_HOST"),
		os.Getenv("MYSQL_TWEAT_PORT"),
		os.Getenv("MYSQL_TWEAT_DATABASE"),
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func TweatGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("MYSQL_TWEAT_USER"),
		os.Getenv("MYSQL_TWEAT_PASSWORD"),
		os.Getenv("MYSQL_TWEAT_HOST"),
		os.Getenv("MYSQL_TWEAT_PORT"),
		os.Getenv("MYSQL_TWEAT_DATABASE"),
	)

	db, err := gorm.Open(gm.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
