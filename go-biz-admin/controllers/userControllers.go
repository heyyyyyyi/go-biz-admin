package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/database"
	"github.com/heyyyyyyi/go-biz-admin/models"
	"github.com/heyyyyyyi/go-biz-admin/utils"
	"net/http"
	"strconv"
)

// 注册，收集前端user信息，交给本函数，写入数据库
func Register(c *gin.Context) {
	var data map[string]string
	// 解析JSON
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "input data is not JSON format"})
		return
	}

	// password 再次确认
	if data["password"] != data["password_confirm"] {
		fmt.Println("password does not match...")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "password does not match"})
		return
	}

	var role_id int
	role_id, _ = strconv.Atoi(data["role_id"])

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		//Password:  data["password"],
		RoleId: uint(role_id),
	}
	//password...
	user.TranslatePassword(data["password"])
	//往表里插入函数
	database.DB.Create(&user)
	//c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "input data is not JSON format"})
		return
	}
	//拉数据库的记录
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	// select * from user where emal = ? limit 1 order by id
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "incorrect password"})
		return
	}

	// jwt
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.SetCookie("jwt", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, user)
}

func User(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	id, _ := utils.ParseJwt(cookie)
	user := models.User{
		Id: uint(id),
	}
	database.DB.Preload("Role").Find(&user)
	c.JSON(http.StatusOK, user)
	return
}

// 观看用户
func AllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Role").Find(&users)
	c.JSON(http.StatusOK, users)
}

// 观看一个user
func FindAUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	user := models.User{
		Id: uint(id),
	}
	database.DB.Preload("Role").Find(&user)
	c.JSON(http.StatusOK, user)
}

// 创建user （管理员）
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"})
		return
	}
	//预设密码
	user.TranslatePassword("1234")
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	user := models.User{
		Id: uint(id),
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"})
		return
	}
	database.DB.Model(&user).Updates(user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	user := models.User{
		Id: uint(id),
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Delete succeed."})
}
