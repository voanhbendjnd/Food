package main

import (
	"FoodDelivery/module/restaurant/transport/ginrestaurant"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (Restaurant) TableName() string       { return "restaurants" }

func main() {
	dsn := os.Getenv("MYSQL_MANAGEMENT")
	//dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=true&Local"
	//dsn := "root:1607@tcp(127.0.0.1:3306)/javasql?charset=utf8mb4&parseTime=true&loc=Local" // workbench

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Running!!!")
	//log.Print(db, err)
	//newRestaurant := Restaurant{Name: "Old Restaurant", Address: "Adam/123"}
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Fatal(err)
	//}
	//var res Restaurant
	//if err := db.Where("id = ?", 4).First(&res).Error; err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(res)
	//var resUpdate Restaurant
	//newName := ""
	//updateName := RestaurantUpdate{Name: &newName}
	//resUpdate.Name = "New Restaurant"
	//if err := db.Where("id = ?", res.Id).Updates(&updateName).Error; err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resUpdate)

	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", res.Id).Delete(nil).Error; err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Delete successfully")

	r := gin.Default()                    // lấy server
	r.GET("/ping", func(c *gin.Context) { // register link on server with /ping
		c.JSON(http.StatusOK, gin.H{
			"message": "Chạy được rồi",
		})
	})
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	//Create
	restaurants.POST("", ginrestaurant.CreateRestaurant(db))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		var data Restaurant
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Id must be a number",
			})
			return
		}
		if db.Where("id = ?", id).First(&data).Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Id restaurant (" + c.Param("id") + ") not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	})
	restaurants.GET("", func(c *gin.Context) {
		var data []Restaurant
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}
		offset := (pagingData.Page - 1) * pagingData.Limit
		db.Offset(offset).
			Order("id desc").
			Limit(pagingData.Limit).
			Find(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	})
	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id must be a number",
			})
			return
		}
		var dataUpdate Restaurant
		if err := c.ShouldBind(&dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var findById Restaurant
		res := db.Where("id = ?", id).Updates(&dataUpdate)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})
			return
		}
		db.Where("id = ?", id).First(&findById)
		c.JSON(http.StatusOK, gin.H{
			"Message": "Update successfully",
			"data":    findById,
		})
	})
	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id must be a number",
			})
			return
		}
		if db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Id restaurant (" + c.Param("id") + ") not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete successfully",
		})
	})
	r.Run(":8080")
}
