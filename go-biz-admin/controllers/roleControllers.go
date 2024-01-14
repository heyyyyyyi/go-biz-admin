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
func AllRoles(c *gin.Context) {
	var roles []models.Role
	database.DB.preload("Role").Find(&roles)
	c.JSON(http.StatusOK, roles)
}

// 观看一个
func FindARole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.preload("Role").Find(&role)
	c.JSON(http.StatusOK, role)
}

// 创建role (permission many2many 关系)
/*
Role JSON
{
	name: "admin"
	permissions: [
		"1",
		"2",
		...
		"8"
	]
}
*/
type RoleCreateDTO struct {
	Name        string   `json: "name"`
	Permissions []string `json: "permissions"`
}

func CreateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var roleDto RoleCreateDTO
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"})
		return
	}
	//预设密码
	permissions := make([]models.Permission, len(roleDto.Permissions))
	for idx, permissionId := range roleDto.Permissions {
		id, _ := strconv.Atoi(permissionId)
		permissions[idx] = models.Permission{Id: uint(id)}
	}
	role := models.Role{
		Name:        roleDto.Name,
		permissions: permissions,
	}
	// ...
	database.DB.Create(&role)
	c.JSON(http.StatusOK, role)
}

/*
Role JSON
{
	"id": "id"
	name: "admin"
	permissions: [
		"1",
		"2",
		...
		"8"
	]
}
*/

func UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var roleDto RoleCreateDTO
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"})
		return
	}
	permissions := make([]models.Permission, len(roleDto.Permissions))
	for idx, permissionId := range roleDto.Permissions {
		id, _ := strconv.Atoi(permissionId)
		permissions[idx] = models.Permission{Id: uint(id)}
	}
	role := models.Role{
		Id:          uint(id),
		Name:        roleDto.Name,
		permissions: permissions,
	}
	// ...
}

/*
Role JSON

	{
		"id": "id"
	}
*/
func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Delete(&role)
	c.JSON(http.StatusOK, gin.H{"message": "Delete succeed."})
}
