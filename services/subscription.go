package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/perqin/shadow-dashboard-go/db"
	"github.com/perqin/shadow-dashboard-go/models"
)

// SubscriptionService provides operations on Subscription
type SubscriptionService interface {
	GetAll(context *gin.Context)
	AddOne(context *gin.Context)
}

type subscriptionServiceImpl struct {
}

func (service subscriptionServiceImpl) GetAll(context *gin.Context) {
	var subscriptions []models.Subscription
	db.Db.Find(&subscriptions)
	context.JSON(http.StatusOK, subscriptions)
}

func (service subscriptionServiceImpl) AddOne(context *gin.Context) {
	var subscription models.Subscription
	if err := context.BindJSON(&subscription); err != nil {
		return
	}
	db.Db.Create(&subscription)
	context.JSON(http.StatusOK, subscription)
}

// NewSubscriptionService returns a new SubscriptionService
func NewSubscriptionService() SubscriptionService {
	return new(subscriptionServiceImpl)
}
