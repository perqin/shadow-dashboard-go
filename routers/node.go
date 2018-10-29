package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/perqin/shadow-dashboard-go/services"
)

func NewNodeGroup(router *gin.Engine) *gin.RouterGroup {
	group := router.Group("/nodes")
	service := services.NewNodeService()
	group.GET("", service.GetAll)
	group.POST("", service.AddOne)
	return group
}
