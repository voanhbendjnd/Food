package main

import (
	appctx "FoodDelivery/component"
	"FoodDelivery/middleware"
	"FoodDelivery/module/restaurant/transport/ginrestaurant"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error lading .env")
	}
	dsn := os.Getenv("MYSQL_MANAGEMENT")
	// dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=true&Local"
	//dsn := "root:1607@tcp(127.0.0.1:3306)/javasql?charset=utf8mb4&parseTime=true&loc=Local" // workbench

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	appContext := appctx.NewAppContext(db)

	r := gin.Default() // lấy server
	db = db.Debug()
	r.Use(middleware.Recover(appContext))
	r.GET("/ping", func(c *gin.Context) { // register link on server with /ping
		c.JSON(http.StatusOK, gin.H{
			"message": "Chạy được rồi",
		})
	})
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	//Create
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

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
	restaurants.GET("", ginrestaurant.FindAll(appContext))
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
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	r.Run(":8080")
}
