package model

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}
type Student struct {
	Id   string  `json:"id" gorm:"column:ID;primary_key;"`
	Name string  `json:"full_name" gorm:"column:FullName;"`
	GPA  float64 `json:"gpa" gorm:"column:GPA;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (Restaurant) TableName() string       { return "restaurants" }

func (Student) TableName() string { return "student" }
