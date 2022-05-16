package lib

import (
	"fmt"
	"github.com/eight-corner/learngo/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=False&loc=Local",
		constants.user, constants.password, constants.host, constants.port, constants.database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open DB")
	}

	if os.Getenv("ENVIRONMENT") != "production" {

	}

	Db = db
}

func CloseDatabase() {
	if Db != nil {
		Db = nil
	}
}
