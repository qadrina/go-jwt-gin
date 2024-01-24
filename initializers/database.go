package initializers

import (
	//"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL") // moved to .ENV "sqlserver://uae:Demo1234@localhost:1434?database=TestDev"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		//log.Fatal("Failed to connect to database")
		panic("Failed to connect to database")
	}
}

// time stamp 5:55
