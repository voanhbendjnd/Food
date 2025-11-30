package restaurantmodel

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}
type RestaurantCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"address"`
}

// func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (RestaurantCreate) TableName() string { return "restaurants" }
