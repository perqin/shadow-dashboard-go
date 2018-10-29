package routers

import "github.com/gin-gonic/gin"

// NewRouter creates new Gin router for the app
func NewRouter() *gin.Engine {
	router := gin.Default()
	NewSubscriptionGroup(router)
	NewNodeGroup(router)
	return router
}
