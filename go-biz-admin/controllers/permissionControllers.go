package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/database"
	"github.com/heyyyyyyi/go-biz-admin/models"
)

// 观看
func AllPermissions(c *gin.Context) {
	var Permissions []models.Permission
	database.DB.preload("Permission").Find(&Permissions)
	c.JSON(http.StatusOK, Permissions)
}

// 观看一个
func FindAPermission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	permission := models.Permission{
		Id: uint(id),
	}
	database.DB.preload("Permission").Find(&permission)
	c.JSON(http.StatusOK, permission)
}
