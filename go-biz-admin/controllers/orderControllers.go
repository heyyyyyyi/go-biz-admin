package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/database"
	"github.com/heyyyyyyi/go-biz-admin/models"
)

//	func AllOrders(c *gin.Context) {
//		var orders []models.Order
//		database.DB.preload("Orders").Find(&orders)
//		c.JSON(http.StatusOK, orders)
//	}

// 采用paginate
// request body includes "page": 1/2
func AllOrders(c *gin.Context) {
	//前端请求体
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	ret := models.Paginate(database.DB, &models.Order{}, page)
	c.JSON(http.StatusOK, ret)
}

// 返回给前端日期 + 销售总额
type Sales struct {
	Date string `json: "date"`
	Sum  string `json: "sum"`
}

func Chart(c *gin.Context) {
	var sales []Sales
	database.DB.Raw(`
		SELECT DATE_FORMAT(orders.created_at, '%Y-%m-%d') as date, SUM(order_items.price * order_items.quantity) as sum
		FROM orders
		JOIN order_items on orders.id = order_items.order_id
		GROUP BY date
	`).Scan(&sales)
	c.JSON(http.StatusOK, sales)
}
