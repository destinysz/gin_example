package router

import (
	"gin_example/internal/controller"

	"github.com/gin-gonic/gin"
)

func RoleRouter(r *gin.Engine) {
	role := controller.NewRole()
	r.GET("/roles/:id", role.GetOne)
	r.GET("/roles", role.GetList)
	r.POST("/roles", role.Create)
	r.DELETE("/roles/:id", role.Delete)
	r.PUT("/roles/:id", role.Update)
}
