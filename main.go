package main

import (
	"errors"
	"log"
	"net/http"
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
type Student struct {
	Id   string  `json:"id" gorm:"column:ID;primary_key;"`
	Name string  `json:"full_name" gorm:"column:FullName;"`
	GPA  float64 `json:"gpa" gorm:"column:GPA;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (Restaurant) TableName() string       { return "restaurants" }

func (Student) TableName() string { return "student" }
func main() {
	//dsn := os.Getenv("MYSQL_MANAGEMENT")
	//dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=true&Local"
	dsn := "root:1607@tcp(127.0.0.1:3306)/javasql?charset=utf8mb4&parseTime=true&loc=Local" // workbench

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
	restaurants.POST("", func(c *gin.Context) {
		var data Restaurant
		if err := c.ShouldBind(&data); err != nil { // request body
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Lỗi rồi bạn ơi!" + err.Error(),
			})
			return
		}
		if db.Create(&data).Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Thua",
			})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"data": data,
			})
		}

	})
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
	students := v1.Group("/students")
	students.POST("", func(c *gin.Context) {
		var data Student
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Add new student fall",
			})
		}
		if db.Create(&data).Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Add new student un success",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	})
	students.PUT("", func(c *gin.Context) {
		var data Student
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Update student un success",
			})
			return
		}
		id := data.Id
		var existById Student
		if err := db.Where("ID = ?", id).First(&existById).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Id student not found",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Lỗi truy vấn cơ sở dữ liệu: " + err.Error(),
				})
			}
			return

		}
		if err := db.Save(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Update student success",
			"data":    data,
		})
	})
	students.GET("/:id", func(c *gin.Context) {
		var data Student
		id := c.Param("id")
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Id student (" + id + ") not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Find by ID",
			"data":    data,
		})

	})
	students.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table(Student{}.TableName()).Where("ID = ?", id).Delete(nil)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}
		if result.RowsAffected == 0 {
			c.Status(404)
			return
		}
		c.Status(204)
	})
	students.GET("", func(c *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.Status(500)
			return
		}
		if pagingData.Limit <= 0 {
			pagingData.Limit = 3
		}
		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		offset := (pagingData.Page - 1) * pagingData.Limit
		var list []Student
		db.Offset(offset).Limit(pagingData.Limit).Find(&list)
		c.JSON(http.StatusOK, gin.H{
			"limit": pagingData.Limit,
			"page":  pagingData.Page,
			"data":  list,
		})

	})

	r.Run(":8080")
}
