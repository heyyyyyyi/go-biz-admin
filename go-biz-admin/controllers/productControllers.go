package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/database"
	"github.com/heyyyyyyi/go-biz-admin/models"
)

func AllProducts(c *gin.Context) {
	var products []models.Product
	database.DB.preload("Products").Find(&products)
	c.JSON(http.StatusOK, products)
}

//...
