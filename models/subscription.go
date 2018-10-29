package models

import (
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	Id    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Nodes []Node
}
