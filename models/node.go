package models

import (
	"github.com/jinzhu/gorm"
)

type Node struct {
	gorm.Model
	Id           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Enabled      bool   `gorm:"default:false" json:"enabled"`
	Type         string `gorm:"default:'ss'" json:"type"`
	LocalAddress string `gorm:"default:'127.0.0.1'" json:"localAddress"`
	LocalPort    int    `json:"localPort"`
	// TODO: Rename JSON field to `serverAddress`
	ServerAddress string `json:"server"`
	ServerPort    int    `json:"serverPort"`
	Password      string `json:"password"`
	Method        string `json:"method"`
	// `ss` specific
	Plugin string `json:"plugin"`
	// `ssr` specific
	Obfs           string       `json:"obfs"`
	ObfsParam      string       `json:"obfsParam"`
	Protocol       string       `json:"protocol"`
	ProtocolParam  string       `json:"protocolParam"`
	Subscription   Subscription `gorm:"foreignkey:SubscriptionId"`
	SubscriptionId int          `json:"subscriptionId"`
}
