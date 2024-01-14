package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/database"
	"github.com/heyyyyyyi/go-biz-admin/models"
)

func AllOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.preload("Orders").Find(&orders)
	c.JSON(http.StatusOK, orders)
}
