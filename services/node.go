package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/perqin/shadow-dashboard-go/db"
	"github.com/perqin/shadow-dashboard-go/models"
)

type NodeService interface {
	GetAll(context *gin.Context)
	AddOne(context *gin.Context)
}

type nodeServiceImpl struct {
}

func (service nodeServiceImpl) GetAll(context *gin.Context) {
	var nodes []models.Node
	db.Db.Find(&nodes)
	context.JSON(http.StatusOK, nodes)
}

func (service nodeServiceImpl) AddOne(context *gin.Context) {
	var node models.Node
	if err := context.BindJSON(&node); err != nil {
		return
	}
	db.Db.Create(&node)
	context.JSON(http.StatusOK, node)
}

func NewNodeService() NodeService {
	return new(nodeServiceImpl)
}
