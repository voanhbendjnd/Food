package restaurantmodel

import (
	"FoodDelivery/common"
	"errors"
	"strings"
	"time"
)

type RestaurantTypeEnum string

const Normal RestaurantTypeEnum = "normal"
const Premium RestaurantTypeEnum = "premium"
const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Id              int       `json:"id" gorm:"column:id;"`
	Name            string    `json:"name" gorm:"column:name;"`
	Address         string    `json:"address" gorm:"column:address;"`
	PhoneNumber     string    `json:"phone_number" gorm:"phone_number;"`
	Rating          int       `json:"rating" gorm:"rating;"`
	CreatedAt       time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"updated_at"`
	Type            string    `json:"type" gorm:"type;"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}
type RestaurantCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"address"`
}

// handle error
func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type RestaurantDTO struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Address     string `json:"address" gorm:"column:address;"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number;"`
}
type ResRestaurant struct {
	Id          int       `json:"id" gorm:"column:id;"`
	Name        string    `json:"name" gorm:"column:name;"`
	Address     string    `json:"address" gorm:"column:address;"`
	PhoneNumber string    `json:"phone_number" gorm:"phone_number;"`
	Rating      int       `json:"rating" gorm:"rating;"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	Type        string    `json:"type" gorm:"type;"`
}

// func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (RestaurantCreate) TableName() string { return "restaurants" }
func (Restaurant) TableName() string       { return "restaurants" }

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
