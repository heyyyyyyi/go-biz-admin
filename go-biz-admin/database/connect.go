package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:patience@/go_biz_admin"), &gorm.Config{})
	if err != nil {
		fmt.Println("database init failed!")
		return
	}
	fmt.Println("database init...", database)
	DB = database
}
