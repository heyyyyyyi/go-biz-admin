package database

import (
	"fmt"
	"github.com/heyyyyyyi/go-biz-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:patience@/go_biz_admin"), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection failed!", db)
		return
	}
	fmt.Println("database init...", db)
	//建表
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	DB = db

}
