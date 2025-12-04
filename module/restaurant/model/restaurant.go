package restaurantmodel

import (
	"FoodDelivery/common"
	"errors"
	"strings"
)

type RestaurantTypeEnum string

const Normal RestaurantTypeEnum = "normal"
const Premium RestaurantTypeEnum = "premium"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	// Id              int    `json:"id" gorm:"column:id;"`
	Name    string             `json:"name" gorm:"column:name;"`
	Address string             `json:"address" gorm:"column:address;"`
	Type    RestaurantTypeEnum `json:"type" gorm:"column:type"`
	// Status          int    `json:"status" gorm:"column:status"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}
type RestaurantCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"address"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

// func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (RestaurantCreate) TableName() string { return "restaurants" }
func (Restaurant) TableName() string       { return "restaurants" }

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
