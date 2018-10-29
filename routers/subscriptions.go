package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/perqin/shadow-dashboard-go/services"
)

// NewSubscriptionGroup creates new Subscription RouterGroup for given router
func NewSubscriptionGroup(router *gin.Engine) *gin.RouterGroup {
	group := router.Group("/subscriptions")
	service := services.NewSubscriptionService()
	group.GET("", service.GetAll)
	group.POST("", service.AddOne)
	return group
}
