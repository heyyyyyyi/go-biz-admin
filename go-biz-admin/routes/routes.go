package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/controllers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	r.GET("/api/users", controllers.AllUsers)
	r.GET("/api/users/:id", controllers.FindAUser)
	r.POST("/api/users", controllers.CreateUser)
	r.POST("/api/users/:id", controllers.UpdateUser)
	r.Delete("/api/users/:id", controllers.DeleteUser)

	r.GET("/api/roles", controllers.AllRoles)
	r.POST("/api/roles", controllers.CreateRole)
	r.GET("/api/roles/:id", controllers.GetRole)
	r.PUT("/api/roles/:id", controllers.UpdateRole)
	r.DELETE("/api/roles/:id", controllers.DeleteRole)

	r.GET("/api/permissions", controllers.AllPermissions)
	return r
}
