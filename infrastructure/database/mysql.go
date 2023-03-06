package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	gm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func GetTweatGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("MYSQL_TWEAT_USER"),
		os.Getenv("MYSQL_TWEAT_PASSWORD"),
		os.Getenv("MYSQL_TWEAT_HOST"),
		os.Getenv("MYSQL_TWEAT_PORT"),
		os.Getenv("MYSQL_TWEAT_DATABASE"),
	)
	if gormDB != nil {
		return gormDB
	}

	db, err := gorm.Open(gm.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB = db

	return gormDB
}
