package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raolootnexii/trasle/model"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, _ = gorm.Open("sqlite3", "trasle.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&model.Song{})
	fmt.Println("Migrated")
}
