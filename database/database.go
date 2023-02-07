package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	databse_url := os.Getenv("DATABASE_URL")
	if databse_url == "" {
		databse_url = "root:Chutiya777@tcp(127.0.0.1:3308)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	}

	dsn := databse_url
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't Connect To Dtabae")
	}
	fmt.Println("Connected to Dtabase")
}
