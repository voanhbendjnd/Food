package usermodel

import (
	"FoodDelivery/common"
	"errors"
	"regexp"
	"strings"
)

type User struct {
	Id      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Email   string `json:"email" gorm:"column:email"`
	Address string `json:"address" gorm:"column:address"`
}

type UserDTO struct {
	Id      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Email   string `json:"email" gorm:"column:email"`
	Address string `json:"address" gorm:"column:address"`
}

func (User) TableName() string {
	return "users"
}

type ResUser struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id" gorm:"column:id"`
	Name            string `json:"name" gorm:"column:name"`
	Email           string `json:"email" gorm:"column:email"`
	Address         string `json:"address" gorm:"column:address"`
}

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func (data *UserDTO) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(data.Email) {
		return ErrEmailIsValidFormat
	}

	return nil
}

const EntityName = "user"

var (
	ErrNameIsEmpty        = errors.New("name cannot be empty")
	ErrEmailIsEmpty       = errors.New("email cannot be empty")
	ErrAddressIsEmpty     = errors.New("address cannot be empty")
	ErrEmailIsValidFormat = errors.New("email invalid")
)
