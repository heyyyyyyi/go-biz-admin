package routes

import (
	"github.com/heyyyyyyi/go-biz-admin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heyyyyyyi/go-biz-admin/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	// + 权限函数 midware
	r.Use(middleware.IsAuthenticate())

	r.GET("/api/users", controllers.AllUsers)
	r.GET("/api/users/:id", controllers.FindAUser)
	r.POST("/api/users", controllers.CreateUser)
	r.POST("/api/users/:id", controllers.UpdateUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)

	r.GET("/api/roles", controllers.AllRoles)
	r.POST("/api/roles", controllers.CreateRole)
	r.GET("/api/roles/:id", controllers.FindARole)
	r.PUT("/api/roles/:id", controllers.UpdateRole)
	r.DELETE("/api/roles/:id", controllers.DeleteRole)

	r.GET("/api/permissions", controllers.AllPermissions)

	r.GET("/api/products", controllers.AllProducts)
	r.POST("/api/products", controllers.CreateProduct)
	r.GET("/api/products/:id", controllers.FindAProduct)
	r.PUT("/api/products/:id", controllers.UpdateProduct)
	r.DELETE("/api/products/:id", controllers.DeleteProduct)

	r.GET("/api/orders", controllers.AllOrders)
	return r
}
