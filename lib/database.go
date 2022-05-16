package lib

import (
	"fmt"
	"github.com/eight-corner/learngo/constants"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=False&loc=Local",
		constants.user, constants.password, constants.host, constants.port, constants.database)

	//db, err :=
}
